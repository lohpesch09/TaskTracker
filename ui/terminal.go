package ui

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Terminal struct {
	userInterface *UserUnterface
}

func NewTerminal(ui *UserUnterface) *Terminal {
	return &Terminal{
		userInterface: ui,
	}
}

func (t *Terminal) GetRequest() error {
	fmt.Println()
	fmt.Print("Введите команду:")

	reader := bufio.NewReader(os.Stdin)

	request, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf(err.Error())
	}

	requestParsed := strings.Fields(request)

	switch requestParsed[0] {
	case "add":
		requestSepBrackets := strings.Split(request, "\"")

		if len(requestSepBrackets) == 3 {
			if requestSepBrackets[2] != "\n" {
				fmt.Println()
				return fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
			}
		} else {
			fmt.Println()
			return fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
		}

		if err = t.userInterface.CreateTask(requestSepBrackets[1]); err != nil {
			fmt.Println()
			return fmt.Errorf(err.Error())
		}
	case "update":
		id, err := strconv.Atoi(requestParsed[1])

		if err != nil {
			return err
		}

		requestSepBrackets := strings.Split(request, "\"")

		if len(requestSepBrackets) == 3 {
			if requestSepBrackets[2] != "\n" {
				fmt.Println()
				return fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
			}
		} else {
			fmt.Println()
			return fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
		}

		if err = t.userInterface.UpdateTask(id, requestSepBrackets[1]); err != nil {
			fmt.Println()
			return fmt.Errorf(err.Error())
		}
	case "delete":
		id, err := strConvToInt(requestParsed)

		if err != nil {
			return fmt.Errorf(err.Error())
		}

		err = t.userInterface.DeleteTask(id)

		if err != nil {
			fmt.Println()
			return fmt.Errorf(err.Error())
		}
	case "mark-in-progress":
		id, err := strConvToInt(requestParsed)

		if err != nil {
			return fmt.Errorf(err.Error())
		}

		err = t.userInterface.MarkTask(id, "в процессе выполнения")

		if err != nil {
			fmt.Println()
			return fmt.Errorf(err.Error())
		}
	case "mark-done":
		id, err := strConvToInt(requestParsed)

		if err != nil {
			return fmt.Errorf(err.Error())
		}

		err = t.userInterface.MarkTask(id, "выполнено")

		if err != nil {
			fmt.Println()
			return fmt.Errorf(err.Error())
		}
	case "list":
		fmt.Println()
		if (len(requestParsed) == 2 && requestParsed[1] != "done" && requestParsed[1] != "todo" && requestParsed[1] != "in-progress") || (len(requestParsed) > 2) {
			return fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
		}

		if len(requestParsed) == 1 {
			if err := t.userInterface.GetAllTasks(); err != nil {
				return err
			}
		} else {
			switch requestParsed[1] {
			case "todo":
				err = t.userInterface.GetTaskDependStatus("не выполнено")
			case "done":
				err = t.userInterface.GetTaskDependStatus("выполнено")
			case "in-progress":
				err = t.userInterface.GetTaskDependStatus("в процессе выполнения")
			}
			if err != nil {
				return err
			}
		}

	case "help":
		fmt.Println()
		if len(requestParsed) > 1 {
			return fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
		}

		comandListFile, err := os.Open("storage/comandsList.txt")

		if err != nil {
			log.Fatalf("Ошибка при попытке открытия файла со списком команд - %v", err)
		}

		comandListReader := bufio.NewReader(comandListFile)

		for {
			line, err := comandListReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("Ошибка при чтении из файла с командами - %v", err)
			}
			fmt.Print(line)
		}

	case "end":
		fmt.Println()
		if len(requestParsed) > 1 {
			return fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
		}

		return errors.New("end")

	default:
		fmt.Println()
		return fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
	}

	return nil
}

func strConvToInt(requestParsed []string) (int, error) {
	if len(requestParsed) > 2 {
		fmt.Println()
		return 0, fmt.Errorf("Вы ввели несуществующую команду. Чтобы посмотреть список доступных команд, введите команду help")
	}

	id, err := strconv.Atoi(requestParsed[1])

	if err != nil {
		fmt.Println()
		return 0, fmt.Errorf(err.Error())
	}

	return id, nil
}
