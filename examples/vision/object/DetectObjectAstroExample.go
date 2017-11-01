package main

import (
	"github.com/rajdeepd/einstein-go-apis/vision"
	"os"
	"encoding/json"
	"fmt"
	"github.com/rajdeepd/einstein-go-apis/common"
)

func main() {
	modelId := "UVNMAD26FSVAVPHOWNFKKN63FM"
	sampleLocation := "https://www.dropbox.com/s/xryf4smcgfih5m0/astro2.JPG?dl=1"
	accessToken := os.Getenv("ACCESSKEY")
	response, err := vision.DetectObject(modelId, sampleLocation, accessToken)
	if err != nil {
		fmt.Println(err)
	}

	var dat *common.PredictResponse
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	for _, value := range dat.Probabilities {
		fmt.Println(value)
	}

}