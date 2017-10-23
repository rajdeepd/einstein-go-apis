package main

import (
	"os"
	"github.com/rajdeepd/einstein-go-apis/vision"
	"fmt"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")
	path := "/home/ubuntu/Downloads/mountainvsbeach.zip"
	response, err := vision.CreateDataSetFromLocalFile(path, accessToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response: " + string(response))
}