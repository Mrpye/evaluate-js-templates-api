package body_types

type Evaluate struct {
	Model interface{} `json:"model"`
	Code  string      `json:"code"`
}

type Template struct {
	Model    interface{} `json:"model"`
	Template string      `json:"template"`
}
