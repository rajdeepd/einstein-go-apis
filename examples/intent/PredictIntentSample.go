package main

import (
	"github.com/rajdeepd/einstein-go-apis/lang"
	"os"
	"encoding/json"
	"fmt"
	"github.com/rajdeepd/einstein-go-apis/common"

)

func main() {
	modelId := "GlobalIntentModel"
	document := "password reset"
	accessToken := os.Getenv("ACCESSKEY")
	response, err := lang.PredictIntent(modelId, document, accessToken)
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