package todo

type UserID int

type TodoID string

type Todo struct {
	ID     TodoID
	UserID UserID
	Text   string
}

type NewTodo struct {
	UserID UserID
	Text   string
}

func (id *TodoID) String() string {
	if id == nil {
		return ""
	}
	return string(*id)
}
