package controller

import "net/http"

func Register() (*http.ServeMux){

	mux := http.NewServeMux()
	mux.HandleFunc("/", ping())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/readAll", readAll())
	mux.HandleFunc("/readByName", readByName())
	mux.HandleFunc("/deleteByName/", deleteByName())
	return mux
}
