package main


import (
	"os"
	"github.com/rajdeepd/einstein-go-apis/vision"
	"fmt"
	"encoding/json"
	"strconv"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")
	datasetId := "1005133"
        name := "Beach and Mountain"
	response, err := vision.TrainDataSet(accessToken, name, datasetId)
	if err != nil {
		fmt.Println(err)
	}
	var dat *vision.TrainDataSetResponse
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	fmt.Println("ModelID: " + dat.ModelID)
	fmt.Println("ModelType: " + dat.ModelType)
	fmt.Println("LearningRate: " + strconv.FormatFloat(dat.LearningRate, 'g', 1, 64))
	fmt.Println("Status: " + dat.Status)
	fmt.Println("QueuePosition: " + strconv.Itoa(dat.QueuePosition))
	fmt.Sprintf("%v,%v","TrainStatus",dat.TrainStats)

}