package main

import (
	"fmt"
	"reflect"
	"sync"
)

var wg sync.WaitGroup

func goroutine1(i int) {
	fmt.Printf("hello goroutine!:%d \n", i)
	wg.Done()
}

type UserInfoStruct struct {
	uid  uint32
	name string
}

func (userinfo UserInfoStruct) getusername() (string, error) {
	if userinfo.uid >= 100 {
		return "bigusername", nil
	} else {
		return "smallusername", nil
	}
}
func (userinfo UserInfoStruct) setusername() (bool, error) {
	if userinfo.uid >= 100 {
		return true, nil
	} else {
		return false, nil
	}
}
func main() {
	//wg.Add(10)
	//for i:=0;i<2;i++ {
	//	go goroutine1(i)
	//}
	//wg.Wait()
	//fmt.Println("main goroutine done!")
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	for k, v := range s {
		fmt.Printf("k=%d\n", k)
		fmt.Printf("v=%d\n", v)
		fmt.Println("===")
	}

	userinfo := UserInfoStruct{
		100, "gencozhang",
	}
	fmt.Println(reflect.TypeOf(userinfo))
	fmt.Println(reflect.ValueOf(userinfo))
	username, err := userinfo.getusername()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(username)

	setflag, err := userinfo.setusername()
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Printf("set flag:%t \n", setflag)

}
