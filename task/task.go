package task

type Task struct {
	description string
	status      string
	createdAt   string
	updatedAt   string
}

func NewTask(description, status, createdAt, updatedAt string) *Task {
	return &Task{
		description: description,
		status:      status,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (task Task) String() string {
	return "Описание задачи: " + task.description + "\n" + "Статус задачи: " + task.status
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) SetStatus(status string) {
	t.status = status
}

func (t *Task) SetUpdatedAt(updatedAt string) {
	t.updatedAt = updatedAt
}

func (t Task) GetDescription() string {
	return t.description
}

func (t Task) GetStatus() string {
	return t.status
}

func (t Task) GetCreatedAt() string {
	return t.createdAt
}

func (t Task) GetUpdatedAt() string {
	return t.updatedAt
}
