package main

import (
	"C"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"service"
	"strings"
	"time"
)

const _TIMEOUT_LIMIT = 60

func call_post_req(apiName string, concurrentOn string, apiData string, client *http.Client, respo chan<- string, quit <-chan bool) {
	data := make(map[string]string)
	data["data"] = apiData
	data["pk"] = concurrentOn
	dataByte, err1 := json.Marshal(data)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	req, err2 := http.NewRequest("POST", apiName, bytes.NewBuffer(dataByte))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	resp, err3 := client.Do(req)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	defer resp.Body.Close()
	body, err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		fmt.Println(err4)
		return
	}
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

func call_grpc_req(apiName string, concurrentOn string, apiData string, client service.InterstallerCallClient, respo chan<- string, quit <-chan bool) {
	data := make(map[string]string)
	data["data"] = apiData
	data["pk"] = concurrentOn
	dataByte, err1 := json.Marshal(data)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	dataString := string(dataByte) 
	response, err2 := client.MakeInterstallerCall(context.Background(), &service.InterstallerRequest{Name: &dataString})
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	select {
	case <-quit:
		return
	case <-time.Tick((1 * time.Millisecond)):
		respo <- *response.Message
	}
}

//export call_rpc
func call_rpc(apiNameC *C.char, stringDataC *C.char, concurrencyOnListC *C.char) *C.char {
	apiName := C.GoString(apiNameC)
	stringData := C.GoString(stringDataC)
	concurrencyOnString := C.GoString(concurrencyOnListC)
	concurrencyOnList := strings.Split(concurrencyOnString, ",")
	concurrentOnCount := len(concurrencyOnList)
	respo := make(chan string, concurrentOnCount)
	quit := make(chan bool, concurrentOnCount)
	conn, err := grpc.Dial(apiName, grpc.WithInsecure())
	if err != nil {
		panic("Can not connect to gRPC.")
		return C.CString(err.Error())
	}
	defer conn.Close()
	client := service.NewInterstallerCallClient(conn, )
	for _, value := range concurrencyOnList {
		go call_grpc_req(apiName, value, stringData, client, respo, quit)
	}
	doneStr := ""
	tick := time.Tick(_TIMEOUT_LIMIT * time.Second)
	flag := true
	for i := 0; flag && i < concurrentOnCount; i++ {
		select {
		case s := <-respo:
			doneStr += s
			fmt.Println(i)
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
	tick := time.Tick(_TIMEOUT_LIMIT * time.Second)
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
