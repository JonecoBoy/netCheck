package event

type TaskCreated struct {
	TaskID string
}

type TaskUpdated struct {
	TaskID string
}

type TaskDeleted struct {
	TaskID string
}
