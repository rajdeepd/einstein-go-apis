package main

import (
	"github.com/rajdeepd/metamind-go-apis/vision"
	"os"
)

func main() {
	modelId := "JNFE34BG6BK6CBYUYOTBL5B2DI"
	samplePath := "/home/ubuntu/Downloads/villa1-test.jpeg"
	accessToken := os.Getenv("ACCESSKEY")
	vision.ClassifyLocal(modelId, samplePath, accessToken)
}