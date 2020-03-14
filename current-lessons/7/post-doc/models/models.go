package models

type List struct {
	Id          int
	Name        string
	Description string
	List        []Task
}

type Task struct {
	Id       int
	Text     string
	Complete bool
}
