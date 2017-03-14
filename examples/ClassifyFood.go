package main

import (
	"github.com/rajdeepd/metamind-go-apis/vision"
	"os"
)

func main() {
	modelId := "FoodImageClassifier"
	sampleLocation := "http://metamind.io/images/foodimage.jpg"
	accessToken := os.Getenv("ACCESSKEY")
	vision.Classify(modelId, sampleLocation, accessToken)
}