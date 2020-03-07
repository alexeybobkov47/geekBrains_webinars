package models

// Описние модели
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}
