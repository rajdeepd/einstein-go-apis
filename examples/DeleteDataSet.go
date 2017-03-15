package main

import (
	"os"
	"github.com/rajdeepd/metamind-go-apis/vision"
	"fmt"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")
	id := "1000681"

	_, err := vision.DeleteDataSet(accessToken, id)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Deleted DataSet: " + id)
	}
}