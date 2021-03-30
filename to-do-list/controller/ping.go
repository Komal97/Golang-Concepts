package controller

import "net/http"

func ping() http.HandlerFunc{

	return func(resp http.ResponseWriter, req *http.Request){
		data := "hello"
		resp.Write([]byte(data))
	}
}
