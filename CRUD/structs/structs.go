package structs

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}
