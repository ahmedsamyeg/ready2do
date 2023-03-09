package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ahmedsamyeg/ready2go/cmd/entity"
	"io"
	"net/http"
)

func Assert(test entity.ApiTest) (bool, error) {
	response, err := http.Get(test.EndPointUrl)
	if err != nil {
		return false, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	mJson, err := json.Marshal(test.ExpectedResponse)
	if err != nil {
		return false, err
	}

	if test.ExpectedStatusCode != 0 && response.StatusCode != test.ExpectedStatusCode {
		message := fmt.Sprintf("expected status %d, but got %d", test.ExpectedStatusCode, response.StatusCode)
		return false, errors.New(message)
	}

	response_match := string(mJson) == string(responseData)

	if !response_match {
		return false, errors.New("Response doesn't match expected response")
	}

	return true, nil
}

//
//	response, err := http.Get(test.EndPoint)
//	responseData, err := io.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	mJson, err := json.Marshal(test.Expects)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(response.StatusCode)
//	fmt.Println(string(responseData))
//	fmt.Println(string(mJson))
//	fmt.Println(string(mJson) == string(responseData))
//	return entity.ApiTestJsonFile{}, nil
//}
