package views

type Respose struct{
	Code int `json:"code"`
	Body interface{} `json:"body"` 
}

type PostRequest struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}
 