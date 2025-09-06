package main

import (
	"fmt"
	"log"
	filework "taskTracker/fileWork"
	"taskTracker/ui"
)

func main() {
	Run()
}

func Run() {
	fileWorker := filework.NewFileWorker()

	taskList, err := fileWorker.Parse()

	if err != nil {
		log.Fatalf("Ошибка чтения данных из файла - %v", err)
	}

	userInterface := ui.NewUserInterface(taskList)

	terminal := ui.NewTerminal(userInterface)

	for {
		err := terminal.GetRequest()
		if err != nil {
			if err.Error() != "end" {
				fmt.Println(err.Error())
			} else if err.Error() == "end" {
				break
			}
		}
	}

	if err = fileWorker.Record(userInterface.GetTaskList()); err != nil {
		log.Fatalf("Ошибка записи данных в файл - %v", err)
	}

	fmt.Println("Завершаю приложение. До свидания!")

}
