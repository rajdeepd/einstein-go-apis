package main

import (
	"os"
	"fmt"
	"github.com/rajdeepd/einstein-go-apis/lang"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")
	path := "/home/ubuntu/work/src/github.com/rajdeepd/einstein-go-apis/data/movie_review_train_1000_v2.tsv"
	response, err := lang.CreateSentimentDataSet(path, accessToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response: " + string(response))
}