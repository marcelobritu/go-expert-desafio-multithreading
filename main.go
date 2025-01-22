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
	go GetAPI("https://brasilapi.com.br/api/cep/v1/"+cep, ch1)
	go GetAPI("http://viacep.com.br/ws/"+cep+"/json", ch2)

	select {
	case result := <-ch1:
		fmt.Printf("BrasilAPI\n---\n%s\n", result)
	case result := <-ch2:
		fmt.Printf("ViaCEP\n---\n%s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
