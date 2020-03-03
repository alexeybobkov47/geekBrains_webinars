package models

type Lists struct {
	Id          int
	Name        string
	Description string
	List        []Task
}

type Task struct {
	Id       int
	ListID   int
	Text     string
	Complete bool
}

type User struct {
	Id   int
	Name string
}
