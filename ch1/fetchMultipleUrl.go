package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/**
fetch多条url的resp  print每条 url的耗时 内容长度 以及url
*/
func main() {
	start := time.Now()
	ch := make(chan string)
	if len(os.Args) > 1 {
		for _, url := range os.Args[1:] {
			go fetch(url, ch)
		}
	} else {
		log.Fatal("no fetch url ")
		os.Exit(1)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	seconds := time.Since(start).Seconds()
	log.Printf("fetch all url consume seconds:%.8f \n", seconds)

}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("fetch error when get,%v\n", err)
		ch <- err.Error()
	}
	fmt.Println(resp.Status)
	lenofbody, err1 := io.Copy(ioutil.Discard, resp.Body)
	fmt.Println(lenofbody)
	defer resp.Body.Close()
	if err1 != nil {
		log.Printf("fetch error when copy,%v\n", err)
		ch <- err1.Error()
	}
	sec := time.Since(start).Seconds()
	//fmt.Println(fmt.Sprintf("%.2fs   %7d    %s ",sec,lenofbody,url))
	ch <- fmt.Sprintf("%.2fs   %7d    %s ", sec, lenofbody, url)

}
