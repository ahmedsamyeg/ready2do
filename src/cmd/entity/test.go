package entity

type Test struct {
	TestName string                 `json:"name,omitempty"`
	EndPoint string                 `json:"endpoint,omitempty"`
	Expects  map[string]interface{} `json:"expects,omitempty"`
}
