package entity

type ApiTest struct {
	Category           string                 `json:"category,omitempty"`
	Title              string                 `json:"title,omitempty"`
	EndPointUrl        string                 `json:"endpoint,omitempty"`
	Header             string                 `json:"header,omitempty"`
	TestType           string                 `json:"test_type,omitempty"`
	Request            map[string]interface{} `json:"request,omitempty"`
	ExpectedResponse   map[string]interface{} `json:"expected_response,omitempty"`
	ExpectedStatusCode int                    `json:"expected_status_code,omitempty"`
}
