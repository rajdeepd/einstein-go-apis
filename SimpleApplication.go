package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)


func main() {
	fmt.Printf("Hello, world.\n")
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		// handle error
	}
	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	fmt.Println(os.Stdout, string(htmlData))
}
