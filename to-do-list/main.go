package main

import (
	"./controller"
	"./model"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main(){

	mux := controller.Register()
	db := model.Connect()

	defer db.Close()
	http.ListenAndServe("localhost:3000", mux)
}
