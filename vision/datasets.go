package vision

import (
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io/ioutil"
)


//curl -X POST -H "Authorization: Bearer <TOKEN>" -H "Cache-Control: no-cache" -H
// "Content-Type: multipart/form-data" -F
// "path=https://dropboxusercontent.com/u/99999999/mountainvsbeach.zip"
// https://api.metamind.io/v1/vision/datasets/upload/sync
func CreateDataSet(path string, accessToken string) ([]byte, error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Println(accessToken)

	_ = writer.WriteField("path", path)
	_ = writer.Close()

	req, err := http.NewRequest("POST", DATASET_UPLOAD_SYNC_URL_V1, body)
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
//curl -X GET -H "Authorization: Bearer <TOKEN>" -H "Cache-Control: no-cache"
// https://api.metamind.io/v1/vision/datasets
func ListDataSets(accessToken string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", DATASET_URL, nil)
	req.Header.Set("Content-Type", "multipart/form-data")

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

//curl -X DELETE -H "Authorization: Bearer <TOKEN>" -H
// "Cache-Control: no-cache" https://api.metamind.io/v1/vision/datasets/<DATASET_ID>

func DeleteDataSet(accessToken, id string) ([]byte, error) {
	//client := &http.Client{}

	req, err := http.NewRequest("DELETE", DATASET_URL + "/" +  id, nil)
	//req.Header.Set("Content-Type", "multipart/form-data")

	// ...
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

//curl -X POST -H "Authorization: Bearer <TOKEN>" -H "Cache-Control: no-cache" -H
// "Content-Type: multipart/form-data" -F "name=Beach Mountain Model" -F
// "datasetId=57" https://api.metamind.io/v1/vision/train

func TrainDataSet(accessToken,name, datasetId string) ([]byte, error) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Println(accessToken)

	_ = writer.WriteField("name", name)
	_ = writer.WriteField("datasetId", datasetId)
	_ = writer.Close()

	req, err := http.NewRequest("POST",
		TRAIN_URL, body)
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

