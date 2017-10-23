package vision

import (
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
	"io"
)


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

func CreateDataSetFromLocalFile(path string, accessToken string) ([]byte, error) {
	client := &http.Client{}
	/*body := &bytes.Buffer{}
	//file_data := getFileData(path)
	writer := multipart.NewWriter(body)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()
	//body := new(bytes.Buffer)
	//writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(path, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	_ = writer.Close()*/
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

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()


	req, err := http.NewRequest("POST", DATASET_UPLOAD_SYNC_URL_V1, &b)
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
//curl -X GET -H "Authorization: Bearer <TOKEN>" -H "Cache-Control: no-cache"
// https://api.einstein.ai/v1/vision/datasets

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


func DeleteDataSet(accessToken, id string) ([]byte, error) {

	req, err := http.NewRequest("DELETE", DATASET_URL + "/" +  id, nil)
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

func getFileData(samplePath string) []byte {
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

	return buf
}
