package main


import (
	"os"
	//"github.com/rajdeepd/einstein-go-apis/vision"
	"fmt"
	"encoding/json"
	//"strconv"
	"github.com/rajdeepd/einstein-go-apis/lang"
	"github.com/rajdeepd/einstein-go-apis/common"
)

func main() {
	accessToken := os.Getenv("ACCESSKEY")
	datasetId := "1016684"
	response, err := lang.GetDataSetDetails(accessToken, datasetId)
	if err != nil {
		fmt.Println(err)
	}
	var dat *common.DataSet
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat.ID)
	fmt.Println(dat.Name)
	fmt.Println(dat.StatusMsg)
	//fmt.Println(dat.LabelSummary)
	fmt.Println(dat.TotalLabels)
	fmt.Println(dat.TotalExamples)

}