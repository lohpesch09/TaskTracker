package ui

import (
	"errors"
	"fmt"
	"taskTracker/task"
	"time"
)

type UserUnterface struct {
	taskList *task.TaskList
}

func NewUserInterface(taskList *task.TaskList) *UserUnterface {
	return &UserUnterface{
		taskList: taskList,
	}
}

func (ui *UserUnterface) GetTaskList() *task.TaskList {
	return ui.taskList
}

func (ui *UserUnterface) CreateTask(description string) error {

	if description == "" {
		return errors.New("вы передали пустое описание задачи! Пожалуйста, введите описание")
	}

	id := len(ui.taskList.GetTaskMap()) + 1

	ui.taskList.SetTaskById(id, task.NewTask(description, "не выполнено", time.Now().Format("2006-01-02 15:04:05"), "не изменялось"))

	return nil
}

func (ui *UserUnterface) UpdateTask(id int, description string) error {
	if description == "" {
		return fmt.Errorf("вы ничего не ввели в качестве описания. Пожалуйста, введите новое описание для задачи %v", id)
	}

	if _, ok := ui.taskList.GetTaskMap()[id]; !ok {
		return fmt.Errorf("задачи с таким id не существует! Попробуйте другой")
	}

	task := ui.taskList.GetTaskById(id)
	task.SetDescription(description)
	task.SetUpdatedAt(time.Now().Format("2006-01-02 15:04:05"))

	ui.taskList.SetTaskById(id, &task)

	return nil
}

func (ui *UserUnterface) DeleteTask(id int) error {
	if _, ok := ui.taskList.GetTaskMap()[id]; !ok {
		return fmt.Errorf("задачи с таким id не существует! Попробуйте другой")
	}

	newTaskMap := ui.taskList.GetTaskMap()

	for i := id; i < len(newTaskMap); i++ {
		newTaskMap[i] = newTaskMap[i+1]
	}

	delete(newTaskMap, len(newTaskMap))

	ui.taskList.SetTaskMap(newTaskMap)

	return nil
}

func (ui *UserUnterface) MarkTask(id int, status string) error {
	if _, ok := ui.taskList.GetTaskMap()[id]; !ok {
		return fmt.Errorf("задачи с таким id не существует! Попробуйте другой")
	}

	newTaskMap := ui.taskList.GetTaskMap()
	task := ui.taskList.GetTaskById(id)

	task.SetStatus(status)

	newTaskMap[id] = task

	return nil
}

func (ui *UserUnterface) GetAllTasks() error {
	ui.taskList.GetTaskMap()

	for i := 1; i <= len(ui.taskList.GetTaskMap()); i++ {
		if i < len(ui.taskList.GetTaskMap()) {
			fmt.Println("Задача id = ", i)
			fmt.Println(ui.taskList.GetTaskMap()[i].String())
			fmt.Println()
		} else {
			fmt.Println("Задача id = ", i)
			fmt.Println(ui.taskList.GetTaskMap()[i].String())
		}

	}

	return nil
}

func (ui *UserUnterface) GetTaskById(id int) error {
	if _, ok := ui.taskList.GetTaskMap()[id]; !ok {
		return fmt.Errorf("задачи с таким id не существует! Попробуйте другой")
	}

	fmt.Println("Задача id = ", id)
	fmt.Println(ui.taskList.GetTaskById(id).String())
	fmt.Println()

	return nil
}

func (ui *UserUnterface) GetTaskDependStatus(status string) error {
	for i := 1; i <= len(ui.taskList.GetTaskMap()); i++ {
		if ui.taskList.GetTaskMap()[i].GetStatus() == status {
			fmt.Println("Задача id = ", i)
			fmt.Println(ui.taskList.GetTaskMap()[i].String())
			fmt.Println()
		}
	}

	return nil
}
