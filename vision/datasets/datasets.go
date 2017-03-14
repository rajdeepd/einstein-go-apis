package main

import (
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	createDataSet()

}
//curl -X POST -H "Authorization: Bearer <TOKEN>" -H "Cache-Control: no-cache" -H
// "Content-Type: multipart/form-data" -F
// "path=https://dropboxusercontent.com/u/99999999/mountainvsbeach.zip"
// https://api.metamind.io/v1/vision/datasets/upload/sync
func createDataSet() {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	path := "https://dropboxusercontent.com/u/99999999/mountainvsbeach.zip"
	accessToken := os.Getenv("ACCESSKEY")
	fmt.Println(accessToken)

	//_ = writer.WriteField("modelId", modelId)
	_ = writer.WriteField("path", path)
	_ = writer.Close()

	req, err := http.NewRequest("POST", "https://api.metamind.io/v1/vision/datasets/upload/sync", body)
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
	htmlData, err := ioutil.ReadAll(resp.Body)
	fmt.Println(os.Stdout, string(htmlData))

}
