package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetAPI(url string, ch chan<- string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	ch <- string(body)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	cep := "60822345"
	url1 := "https://brasilapi.com.br/api/cep/v1/" + cep
	url2 := "http://viacep.com.br/ws/" + cep + "/json"

	go GetAPI(url1, ch1)
	go GetAPI(url2, ch2)

	select {
	case result := <-ch1:
		fmt.Printf("%s\n-------------------------------\n%s\n", url1, result)
	case result := <-ch2:
		fmt.Printf("%s\n-------------------------------\n%s\n", url2, result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
