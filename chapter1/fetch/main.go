package main

import (
	"fmt"
	//"io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		ok := strings.HasPrefix(url, "http://")
		if !ok {
			url = "http://" + url
		}
		fmt.Printf("accessing %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("Status Code: %s\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}

	}

}
