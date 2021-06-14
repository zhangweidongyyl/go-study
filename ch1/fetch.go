package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Println(resp.Status)

		b, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			log.Fatal(err1)
			os.Exit(1)
		}
		log.Println(string(b))
		defer resp.Body.Close()
	} else {
		log.Println("please enter fetch url")
		os.Exit(0)
	}

}
