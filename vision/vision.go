package vision

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"mime/multipart"
	"bytes"
	"encoding/base64"

	"bufio"
)


func Classify(modelId string, sampleLocation string, accessToken string){
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("modelId", modelId)
	_ = writer.WriteField("sampleLocation", sampleLocation)
	_ = writer.Close()

	req, err := http.NewRequest("POST", "https://api.metamind.io/v1/vision/predict", body)
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

func ClassifyLocal(modelId string, samplePath string, accessToken string){
	base64 := getBase64String(samplePath)
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("modelId", modelId)
	_ = writer.WriteField("sampleBase64Content", base64)
	_ = writer.Close()

	req, err := http.NewRequest("POST", "https://api.metamind.io/v1/vision/predict", body)
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

func getBase64String(samplePath string) string {
	imgFile, err := os.Open(samplePath)

	if err != nil {
	fmt.Println(err)
	  os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)


	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return imgBase64Str
}


