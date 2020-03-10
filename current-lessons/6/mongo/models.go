package main

type List struct {
	ID          int
	Name        string
	Description string
	List        []Task
}

type Task struct {
	Id       int
	Text     string
	Complete bool
}
