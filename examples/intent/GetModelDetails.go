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
	modelId := "BMRG2U4ST6UP3OZAO2J7N5MF5M"
	response, err := lang.GetModelDetails(accessToken, modelId)
	if err != nil {
		fmt.Println(err)
	}
	var model *common.Model
	if err := json.Unmarshal(response, &model); err != nil {
		panic(err)
	}
	fmt.Println(model.ID)
	fmt.Println(model.MetricsData)
	/*fmt.Println(dat.Name)
	fmt.Println(dat.StatusMsg)
	//fmt.Println(dat.LabelSummary)
	fmt.Println(dat.TotalLabels)
	fmt.Println(dat.TotalExamples)*/

}