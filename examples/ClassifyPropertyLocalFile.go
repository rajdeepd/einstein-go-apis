package main

import (
	"github.com/rajdeepd/einstein-go-apis/vision"
	"os"
	"fmt"
	"encoding/json"
)

func main() {
	modelId := "JNFE34BG6BK6CBYUYOTBL5B2DI"
	samplePath := "/home/ubuntu/Downloads/villa1-test.jpeg"
	accessToken := os.Getenv("ACCESSKEY")
	response, err := vision.ClassifyLocal(modelId, samplePath, accessToken)
	if err != nil {
		fmt.Println(err)
	}

	var dat *vision.PredictResponse
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	for _, value := range dat.Probabilities {
		fmt.Println(value)
	}
}