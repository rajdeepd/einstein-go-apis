package lang

import (
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io/ioutil"
	//"github.com/rajdeepd/einstein-go-apis/common"
)
//  builder.addTextBody("document", document)
//builder.addTextBody("modelId", modelId)

func PredictIntent(modelId string, document string, accessToken string) ([]byte, error){
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("modelId", modelId)
	_ = writer.WriteField("document", document)
	_ = writer.Close()

	req, err := http.NewRequest("POST", BASE_INTENT, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// ...
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "Bearer " +  accessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	respBody, errResp := ioutil.ReadAll(resp.Body)
	return respBody, errResp
}

