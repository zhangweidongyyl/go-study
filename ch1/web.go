package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var counter uint32

func main() {
	web()
}
func web() {
	http.HandleFunc("/", handler)
	//http.HandleFunc("/abc",handlerAbc)
	http.HandleFunc("/count", countreq)
	http.ListenAndServe(":8098", nil)
}
func handler(rw http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++

	rw.Write([]byte("hello 1"))
	mu.Unlock()
	fmt.Println(counter)
}

//
//func handlerAbc(rw http.ResponseWriter,r *http.Request) {
//	mu.Lock()
//	counter ++
//	rw.Write([]byte("hello 2"))
//	mu.Unlock()
//}
func countreq(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(fmt.Sprintf("counter:%d", counter)))
}
