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
/*
extern char* call_api(char *, char *, char *);

static char* do_call_api(char* a, char* b, char *c) {
	return call_api(a, b, c);

*/
func call_post_req(apiName string, concurrentOn string, apiData string, client *http.Client, respo chan<- string, quit <-chan bool) {
	data := make(map[string]string)
	data["data"] = apiData
	data["pk"] = concurrentOn
	dataByte, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", apiName, bytes.NewBuffer(dataByte))
	//req, _ := http.NewRequest("POST", "http://localhost:8000/policy/v1/get-policy/", bytes.NewBuffer(dataByte))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal([]byte(body), &response)
	fmt.Println(string(body))
	fmt.Println(concurrentOn)
	// respo <- string(body)
	select{
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
	fmt.Println(apiName, stringData, concurrencyOnString)
	concurrencyOnList := strings.Split(concurrencyOnString, ",")
	respo := make(chan string, len(concurrencyOnList))
	quit := make(chan bool)
	fmt.Println(apiName)
	fmt.Println(stringData)
	fmt.Println(concurrencyOnString)
	client := &http.Client{}
	for _, value := range concurrencyOnList {
		go call_post_req(apiName, value, stringData, client, respo, quit)
		//call_post_req(apiName, value, stringData, client, respo)
	}
	doneStr := ""
	tick := time.Tick(10 * time.Second)
	flag := true
	for flag {
		select{
		case s := <-respo:
			doneStr +=s
		case <-tick:
			fmt.Println("Timeout")
			doneStr += "--Timeout--"
			flag = false
			quit <- true
			break
		}
	}
	// for i := 0; i < len(concurrencyOnList); i++ {
	// 	fmt.Println("waiting for channel")
	// 	doneStr = <-respo
	// 	// fmt.Println(<-respo)
	// }
	close(respo)
	defer close(quit)
	fmt.Println("GO END")
	// doneStr := "We re done hhheheheheheheheheheheh he hebfreifirb ierf 2riu giejrn ijwern kwejng"
	return C.CString(doneStr)
	// return doneStr
}

func main() {}
