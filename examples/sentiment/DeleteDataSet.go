package main

import (
	"os"
	"github.com/rajdeepd/einstein-go-apis/lang"
	"fmt"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")
	id := "1016691"

	_, err := lang.DeleteDataSet(accessToken, id)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Deleted DataSet: " + id)
	}
}