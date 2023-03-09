package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmedsamyeg/ready2go/cmd/entity"
	"github.com/ahmedsamyeg/ready2go/cmd/service"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		printHelp()
		return
	}

	file, err := os.ReadFile(arguments[0])
	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(1)
	}

	for i := 0; i < 100000; i++ {
		tests, err := service.TestFileParser{}.Parse(&file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		Process(&tests)
	}
}

func Process(tests *entity.ApiTestJsonFile) {
	for _, t := range tests.Tests {
		pass, err := service.Assert(t)
		if err != nil || !pass {
			Fail(t, err)
			continue
		}
		Pass(t)
	}
}

func Pass(t entity.ApiTest) {
	message := fmt.Sprintf("[ PASS ] %s - %s", t.Category, t.Title)
	fmt.Println(message)
}

func Fail(t entity.ApiTest, err error) {
	message := fmt.Sprintf("[ FAIL ] %s - %s: %s", t.Category, t.Title, err.Error())
	fmt.Println(message)
}

func printHelp() {
	fmt.Println("Ready2Go")
	fmt.Println("Api test automation tool.")
	fmt.Println("USAGE: readytogo <FILENAME>")
}

func TestSingleEndPoint(err error, file []byte) {
	var test entity.Test
	err = json.Unmarshal(file, &test)
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Get(test.EndPoint)
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	mJson, err := json.Marshal(test.Expects)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(string(responseData))
	fmt.Println(string(mJson))
	fmt.Println(string(mJson) == string(responseData))
}
