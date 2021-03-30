package controller

import (
	"../model"
	"../view"
	"encoding/json"
	"fmt"
	"net/http"
)

func create() http.HandlerFunc{
	return func(resp http.ResponseWriter, req *http.Request){
		if req.Method == http.MethodPost{
			// take some data and save it

			data := view.ToDo{}
			_ = json.NewDecoder(req.Body).Decode(&data)			// accept data in body

			err := model.CreateToDo(data.Name, data.ToDo)

			if err != nil{
				resp.WriteHeader(http.StatusInternalServerError)
				resp.Write([]byte(err.Error()))
				fmt.Println("Error in create controller: ", err)
				return
			}

			fmt.Println("Data saved: ")
			resp.WriteHeader(http.StatusCreated)
			json.NewEncoder(resp).Encode(data)

		}
	}
}

func readAll() http.HandlerFunc{
	return func(resp http.ResponseWriter, req *http.Request){
		if req.Method == http.MethodGet{

			data, err := model.ReadAll()
			if err != nil{
				resp.WriteHeader(http.StatusInternalServerError)
				resp.Write([]byte(err.Error()))
				fmt.Println("Error in read all controller: ", err)
				return
			}
			resp.WriteHeader(http.StatusOK)
			json.NewEncoder(resp).Encode(data)
		}
	}
}

func readByName() http.HandlerFunc{
	return func(resp http.ResponseWriter, req *http.Request){
		if req.Method == http.MethodGet{

			name := req.URL.Query().Get("name")		// accept query params
			data, err := model.ReadByName(name)
			if err != nil{
				resp.WriteHeader(http.StatusInternalServerError)
				resp.Write([]byte(err.Error()))
				fmt.Println("Error in read by name controller: ", err)
				return
			}
			resp.WriteHeader(http.StatusOK)
			json.NewEncoder(resp).Encode(data)
		}
	}
}

func deleteByName() http.HandlerFunc{
	return func(resp http.ResponseWriter, req *http.Request){

		if req.Method == http.MethodDelete{
			name := req.URL.Path[1:]							// accept path params
			err := model.DeleteByName(name)
			if err != nil{
				resp.WriteHeader(http.StatusInternalServerError)
				resp.Write([]byte(err.Error()))
				fmt.Println("Error in delete by name controller: ", err)
				return
			}
			resp.WriteHeader(http.StatusOK)
			json.NewEncoder(resp).Encode(struct {
				Status string `json: "status"`
			}{"Item deleted"})
		}
	}
}
