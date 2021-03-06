package views

type Response struct {
	Code int `json:"code"`
	Body interface{} `json:"body"`
}

type PostRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type GetRequest struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}
