package todo

type TodoID string

type Todo struct {
	ID   TodoID
	Text string
}

type NewTodo struct {
	Text string
}

func (id *TodoID) String() string {
	if id == nil {
		return ""
	}
	return string(*id)
}