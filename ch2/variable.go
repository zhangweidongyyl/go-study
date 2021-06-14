package main

import "fmt"

/**
指针 地址等变量
*/
func main() {
	var i int32 = 3

	p := &i

	*p = 333

	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(f() == f())

	incrtest := 99
	incr(&incrtest)
	fmt.Println(incrtest)
}
func incr(p *int) int {
	*p++
	return *p
}
func f() *int {
	v := 4
	return &v
}
