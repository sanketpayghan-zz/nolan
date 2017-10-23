package main

import (
	"C"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func call_post_req(apiName string, concurrentOn string, apiData string, client *http.Client, respo chan<- string, quit <-chan bool) {
	data := make(map[string]string)
	data["data"] = apiData
	data["pk"] = concurrentOn
	dataByte, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", apiName, bytes.NewBuffer(dataByte))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal([]byte(body), &response)
	select {
	case <-quit:
		return
	case <-time.Tick((1 * time.Millisecond)):
		respo <- string(body)
	}
	return
}

//export call_api
func call_api(apiNameC *C.char, stringDataC *C.char, concurrencyOnListC *C.char) *C.char {
	apiName := C.GoString(apiNameC)
	stringData := C.GoString(stringDataC)
	concurrencyOnString := C.GoString(concurrencyOnListC)
	concurrencyOnList := strings.Split(concurrencyOnString, ",")
	concurrentOnCount := len(concurrencyOnList)
	respo := make(chan string, concurrentOnCount)
	quit := make(chan bool, concurrentOnCount)
	client := &http.Client{Timeout: time.Second * 10}
	for _, value := range concurrencyOnList {
		go call_post_req(apiName, value, stringData, client, respo, quit)
		//call_post_req(apiName, value, stringData, client, respo)
	}
	doneStr := ""
	tick := time.Tick(10 * time.Second)
	flag := true
	for i := 0; flag && i < concurrentOnCount; i++ {
		select {
		case s := <-respo:
			doneStr += s
		case <-tick:
			for i := 0; i < concurrentOnCount; i++ {
				quit <- true
			}
			fmt.Println("Timeout")
			doneStr += "--Timeout--"
			flag = false
			break
		}
	}
	close(respo)
	defer close(quit)
	fmt.Println("GO END")
	return C.CString(doneStr)
}

func main() {}
