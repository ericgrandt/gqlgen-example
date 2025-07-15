package model

type TodoTag struct {
	ID     string `json:"id"`
	TodoID string `json:"todo_id"`
	TagID  string `json:"tag_id"`
}

type NewTodoTag struct {
	TodoID string `json:"todo_id"`
	TagID  string `json:"tag_id"`
}
