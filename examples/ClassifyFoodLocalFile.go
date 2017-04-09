package main

import (
	"github.com/rajdeepd/einstein-go-apis/vision"
	"os"
)

func main() {
	modelId := "FoodImageClassifier"
	samplePath := "/home/ubuntu/Downloads/foodimage.jpg"
	accessToken := os.Getenv("ACCESSKEY")
	vision.ClassifyLocal(modelId, samplePath, accessToken)
}