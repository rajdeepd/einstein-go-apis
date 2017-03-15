package main

import (
	"os"
	"github.com/rajdeepd/metamind-go-apis/vision"
	"fmt"
	"encoding/json"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")

	response, err := vision.ListDataSets(accessToken)
	if err != nil {
		fmt.Println(err)
	}

	var dat *vision.DataSetList
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	for _, value := range dat.DataSet {
		fmt.Println(value)
	}
}