package web

type Response struct {
	Code int `json:"code"`
	Message string `json:"message,omitempty"`
	Items []interface{} `json:"items,omitempty"`
}
