package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
)

// custom write implementing 'Writer Interface'
type logWriter struct{}

func (logWriter) Write (bs []byte) (int, error){

	fmt.Println(string(bs))
	return len(bs), nil
}

func getResponse(){

	// resp = *Response and err = error type
	resp, err := http.Get("http://www.google.com")
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("Response type: ", reflect.TypeOf(resp))
	fmt.Println("Response data: ", resp)

	// method - 1
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	// method - 2
	 io.Copy(os.Stdout, resp.Body)

	// method - 3 using custom Writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func implementHttp(){

	fmt.Println("--------------- HTTP -----------------")
	getResponse()
}


