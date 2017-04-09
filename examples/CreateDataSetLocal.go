package main

import (
	"os"
	"github.com/rajdeepd/einstein-go-apis/vision"
	"fmt"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")
	path := "https://s3.ap-south-1.amazonaws.com/rd-einstein-vision/villasvsapartment.zip"
	response, err := vision.CreateDataSet(path, accessToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response: " + string(response))
}