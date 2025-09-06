package filework

import (
	"encoding/json"
	"os"
	"taskTracker/task"
)

const filePath string = "storage/taskTracker.json"

type FileWorker struct{}

func NewFileWorker() *FileWorker {
	return &FileWorker{}
}

func (fw *FileWorker) Parse() (*task.TaskList, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, 0644)

	if err != nil {
		return task.NewTaskList(), err
	}

	fileInfo, err := file.Stat()

	if err != nil {
		return task.NewTaskList(), err
	}

	fileContents := make([]byte, fileInfo.Size())

	_, err = file.Read(fileContents)

	if err != nil {
		return task.NewTaskList(), err
	}

	taskList := task.NewTaskList()

	if len(fileContents) != 0 {
		err = customUnmarshal(&fileContents, taskList)

		if err != nil {
			return task.NewTaskList(), err
		}
	}

	return taskList, nil
}

func (fw *FileWorker) Record(taskList *task.TaskList) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	var taskListContents []byte

	taskListContents, err = customMarshal(taskList)

	if err != nil {
		return err
	}

	_, err = file.Write(taskListContents)

	if err != nil {
		return err
	}

	return nil
}

func customMarshal(taskList *task.TaskList) ([]byte, error) {
	type task struct {
		Description string
		Status      string
		CreatedAt   string
		UpdatedAt   string
	}

	mapForMarshal := make(map[int]task, len(taskList.GetTaskMap()))

	for k, v := range taskList.GetTaskMap() {
		mapForMarshal[k] = task{
			Description: v.GetDescription(),
			Status:      v.GetStatus(),
			CreatedAt:   v.GetCreatedAt(),
			UpdatedAt:   v.GetUpdatedAt(),
		}
	}

	slice, err := json.Marshal(struct {
		TaskMap map[int]task
	}{
		TaskMap: mapForMarshal,
	})

	return slice, err
}

func customUnmarshal(fileContents *[]byte, taskList *task.TaskList) error {
	type Task struct {
		Description string
		Status      string
		CreatedAt   string
		UpdatedAt   string
	}

	var temp struct {
		TaskMap map[int]Task
	}

	if err := json.Unmarshal(*fileContents, &temp); err != nil {
		return err
	}

	mapForUnmarshal := make(map[int]task.Task, len(temp.TaskMap))

	for k, v := range temp.TaskMap {
		mapForUnmarshal[k] = *task.NewTask(v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
	}

	taskList.SetTaskMap(mapForUnmarshal)

	return nil
}
