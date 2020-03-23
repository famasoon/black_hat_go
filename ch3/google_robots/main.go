package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Fatalln("Unable to accept connection")
	}

	fmt.Println("Status: " + resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
