package model

import (
	"../view"
	"fmt"
)

func CreateToDo(name, todo string) error{

	query, err := con.Query("INSERT INTO TO_DO VALUES (?, ?)", name, todo)
	defer query.Close()

	if err != nil{
		return err
	}
	return nil
}

func ReadAll() ([]view.ToDo, error){

	rows, err := con.Query("Select * from To_do")
	if err != nil{
		return nil, err
	}

	fmt.Println("rows: ", rows)
	var todos []view.ToDo
	for rows.Next(){
		data := view.ToDo{}
		rows.Scan(&data.Name, &data.ToDo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]view.ToDo, error){

	rows, err := con.Query("Select * from To_do where name = ?", name)
	if err != nil{
		return nil, err
	}

	fmt.Println("rows: ", rows)
	var todos []view.ToDo
	for rows.Next(){
		data := view.ToDo{}
		rows.Scan(&data.Name, &data.ToDo)
		todos = append(todos, data)
	}
	return todos, nil
}


func DeleteByName(name string) error{

	query, err := con.Query("Delete from To_do where name = ?", name)
	defer query.Close()

	if err != nil{
		return err
	}
	return nil
}
