package lang

import (
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io/ioutil"
	"io"
)



func CreateIntentDataSet(path string, accessToken string) ([]byte, error) {
	return CreateDataSet(path, accessToken, "text-intent")
}

func CreateSentimentDataSet(path string, accessToken string) ([]byte, error) {
	return CreateDataSet(path, accessToken, "text-sentiment")

}

func CreateDataSet(path string, accessToken string, langType string) ([]byte, error) {
	client := &http.Client{}

	fmt.Println(accessToken)
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(path)
	if err != nil {
		return nil,err
	}
	defer f.Close()
	fw, err := w.CreateFormFile("data", path)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return nil, err
	}

	_ = w.WriteField("type", langType)
	w.Close()


	req, err := http.NewRequest("POST", DATASETS_LANG_UPLOAD, &b)
	req.Header.Set("Content-Type", w.FormDataContentType())

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

func GetDataSetDetails(accessToken string, datasetId string) ([]byte, error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Println(accessToken)

	_ = writer.Close()

	req, err := http.NewRequest("GET",
		DATASETS + "/" + datasetId, body)
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

func GetModelDetails(accessToken string, modelId string) ([]byte, error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Println(accessToken)

	_ = writer.Close()

	req, err := http.NewRequest("GET",
		MODELS + "/" + modelId, body)
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



func TrainDataSet(accessToken,name, datasetId string) ([]byte, error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Println(accessToken)

	_ = writer.WriteField("name", name)
	_ = writer.WriteField("datasetId", datasetId)
	_ = writer.Close()

	req, err := http.NewRequest("POST",
		LANG_TRAIN_URL, body)
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

func DeleteDataSet(accessToken, id string) ([]byte, error) {

	req, err := http.NewRequest("DELETE", LANG_TRAIN_URL + "/" +  id, nil)
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "Bearer " +  accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	respBody, errResp := ioutil.ReadAll(resp.Body)
	return respBody, errResp
}

