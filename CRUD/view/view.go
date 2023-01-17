package view

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type PostRequest struct {
	Name string `json:"name"`
	ToDo string `json:"todo"`
}
