package todo

type TodoID string
type Todo struct {
	ID     TodoID
	Text   string
	UserID int32
}

type NewTodo struct {
	Text   string
	UserID int32
}

func (id *TodoID) String() string {
	if id == nil {
		return ""
	}
	return string(*id)
}
