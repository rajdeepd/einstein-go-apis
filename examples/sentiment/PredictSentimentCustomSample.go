package main

import (
	"github.com/rajdeepd/einstein-go-apis/lang"
	"os"
	"encoding/json"
	"fmt"
	"github.com/rajdeepd/einstein-go-apis/common"

)

func main() {
	modelId := "BMRG2U4ST6UP3OZAO2J7N5MF5M"
	document := "this movie is great"
	accessToken := os.Getenv("ACCESSKEY")
	response, err := lang.PredictSentiment(modelId, document, accessToken)
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