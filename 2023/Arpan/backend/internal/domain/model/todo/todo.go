package todo

type TodoID string
type Todo struct {
	ID   TodoID
	User *User
	Text string
}

type NewTodo struct {
	User *User
	Text string
}

func (id *TodoID) String() string {
	if id == nil {
		return ""
	}
	return string(*id)
}

type User struct {
	ID   string
	Name string
}
