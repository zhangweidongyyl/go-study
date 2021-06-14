package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/**
  外部输入文件统计和内部控制台输入
*/
func main() {
	filepath := "D://a.txt"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
		return
	}
	dataarr := strings.Split(string(data), "\n")
	for k, v := range dataarr {
		log.Printf("k=%d,v=%s\n", k, v)
	}
	log.Println()

	counts := make(map[string]uint32)

	if len(os.Args) <= 1 {
		CountLines(os.Stdin, counts)
	} else {
		filepaths := os.Args[1:]
		for _, filepath := range filepaths {
			file, err := os.Open(filepath)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer file.Close()
			CountLines(file, counts)
		}
	}

	for k, v := range counts {
		fmt.Printf("统计结果为：%s出现了%d次\n", k, v)
	}
}

func CountLines(file *os.File, counts map[string]uint32) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputtext := scanner.Text()
		if inputtext == "exit" {
			break
		}
		counts[inputtext]++
	}
}
