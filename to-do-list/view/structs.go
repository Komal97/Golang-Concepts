package view

type Response struct{
	Code int `json: "code"`
	Body interface{} `json: "body"`
}

type ToDo struct {
	Name string `json: "name"`
	ToDo string `json: "todo"`
}
