package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"sync"
)

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return string(body)
}

func TestClient(t *testing.T) {
	time.Sleep(time.Second)
	go func() {
		fmt.Println(httpGet("http://127.0.0.1:8080"))
	}()

	time.Sleep(time.Second)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(httpGet("http://127.0.0.1:8081"))
			wg.Done()
		}()
	}
	wg.Wait()
}
