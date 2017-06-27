package lang

import (
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io/ioutil"
)



func CreateIntentDataSet(path string, accessToken string) ([]byte, error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Println(accessToken)

	_ = writer.WriteField("path", path)
	_ = writer.WriteField("type", "text-intent")

	_ = writer.Close()

	req, err := http.NewRequest("POST", DATASETS_INTENT_UPLOAD, body)
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

