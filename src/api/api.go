
package api

import (
	"net/http"
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"
)

func _makeRequest(method, url string, data map[string]string, client *http.Client) (map[string]interface{}, []byte, error) {
	var response map[string]interface{}
	var b []byte
	dataByte, err1 := json.Marshal(data)
	if err1 != nil {
		fmt.Println(err1)
		return response, b, err1
	}
	req, err2 := http.NewRequest(method, url, bytes.NewBuffer(dataByte))
	if err2 != nil {
		return response, b, err2
	}
	if method == "GET" {
		query := req.URL.Query()
		for key, value := range data {
			query.Add(key, value)
		}
	}
	resp, err3 := client.Do(req)
	if err3 != nil {
		return response, b, err3
	}
	defer resp.Body.Close()
	body, err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		return response, b, err4
	}
	json.Unmarshal([]byte(body), &response)
	return response, body, nil
}

func Post(url string, data map[string]string, client *http.Client, respo chan<- string, quit <-chan bool) {
	_ , body, err := _makeRequest("POST", url, data, client)
	if err != nil {
		fmt.Println(err)
		return
	}
	select {
	case <-quit:
		return
	case <-time.Tick((1 * time.Millisecond)):
		respo <- string(body)
	}
	return
}

func Get(url string, data map[string]string, client *http.Client, respo chan<- string, quit <-chan bool) {
	_ , body, err := _makeRequest("GET", url, data, client)
	if err != nil {
		fmt.Println(err)
		return
	}
	select {
	case <-quit:
		return
	case <-time.Tick((1 * time.Millisecond)):
		respo <- string(body)
	}
	return
}
