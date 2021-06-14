package main

import (
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("./study.log", os.O_CREATE|os.O_APPEND, 0666)
	// 没有则创建，有则追加
	if err != nil {
	}
	log.SetOutput(file)
	log.SetPrefix("gencozhang:")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}
func main() {
	log.Println("s")
}
func logger() {
	log.Println()
}
