package main

import (
	"database/sql"
	"fmt"
	"log"
)

// getAllLists — получение всех списков с задачами
func getAllLists(db *sql.DB) ([]TaskList, error) {
	res := make([]TaskList, 0, 1)

	rows, err := db.Query("select * from task_list_app.lists")
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		list := TaskList{}

		if err := rows.Scan(&list.ID, &list.Name, &list.Description); err != nil {
			log.Println(err)
			continue
		}

		res = append(res, list)
	}

	return res, nil
}

// getList — получение списка по id
func getList(db *sql.DB, id string) (TaskList, error) {
	list := TaskList{}

	row := db.QueryRow(fmt.Sprintf("select * from task_list_app.lists where lists.id = %v", id))
	err := row.Scan(&list.ID, &list.Name, &list.Description)
	if err != nil {
		return list, err
	}

	rows, err := db.Query(fmt.Sprintf("select * from task_list_app.tasks WHERE tasks.list_id = %v", id))
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, new(int), &task.Text, &task.Complete)
		if err != nil {
			log.Println(err)
			continue
		}

		list.List = append(list.List, task)
	}

	return list, nil
}
