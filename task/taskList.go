package task

type TaskList struct {
	taskMap map[int]Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		taskMap: make(map[int]Task, 10),
	}
}

func (taskList *TaskList) SetTaskMap(taskMap map[int]Task) {
	taskList.taskMap = taskMap
}

func (taskList TaskList) GetTaskMap() map[int]Task {
	return taskList.taskMap
}

func (taskList *TaskList) GetTaskById(id int) Task {
	return taskList.taskMap[id]
}

func (taskList *TaskList) SetTaskById(id int, task *Task) {
	taskList.taskMap[id] = *task
}
