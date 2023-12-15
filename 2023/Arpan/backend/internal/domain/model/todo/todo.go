package todo

type TodoID string

type Todo struct {
	ID   TodoID
	Text string
}

type NewTodo struct {
	Text string
}
