package goutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func HttpGetJson(url string, resp interface{}) error {
	result, err := http.Get(url)
	if err == nil {
		defer result.Body.Close()
		return json.NewDecoder(result.Body).Decode(resp)
	}
	return err
}

func HttpGetJsonFromHeader(url string, header map[string]string, resp interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	result, err := client.Do(req)
	if err != nil {
		return err
	}
	defer result.Body.Close()
	return json.NewDecoder(result.Body).Decode(resp)
}

func HttpPostJson(url string, header map[string]string, data map[string]interface{}, resp interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return HttpPostData(url, header, b, resp)
}

func HttpPostData(url string, header map[string]string, b []byte, resp interface{}) error {
	reader := bytes.NewReader(b)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json;charset=UTF-8")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for k, v := range header {
		request.Header.Set(k, v)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(resp)
}

func PostFile(filename string, data map[string]string, targetUrl string) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for k, v := range data {
		if err := bodyWriter.WriteField(k, v); err != nil {
			fmt.Println("write filed error, k:", k, ", v:", v)
		}
	}

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return []byte{}, err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return []byte{}, err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return []byte{}, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
