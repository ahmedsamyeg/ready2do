package entity

type ApiTestJsonFile struct {
	Version     string    `json:"version,omitempty"`
	ProjectName string    `json:"project_name,omitempty"`
	Environment string    `json:"environment,omitempty"`
	Tests       []ApiTest `json:"tests,omitempty"`
}
