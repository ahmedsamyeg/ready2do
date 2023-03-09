package service

import (
	"encoding/json"
	"github.com/ahmedsamyeg/ready2go/cmd/entity"
)

type TestFileParser struct {
}

func (p TestFileParser) Parse(file *[]byte) (entity.ApiTestJsonFile, error) {
	var test entity.ApiTestJsonFile
	err := json.Unmarshal(*file, &test)

	if err != nil {
		return entity.ApiTestJsonFile{}, err
	}

	return test, nil
}
