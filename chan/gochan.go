package main

import "fmt"

func main() {

	//nocachechan()
	//cachechan()
	sendtochan()
}

func nocachechan() {
	// 无缓冲通道 直接发送数据而无接收方时会出现panic
	//var c chan int
	c := make(chan int)
	// 使用make而不是使用var创建，var创建一个nil 通道
	go recv(c)
	c <- 10
	fmt.Println("向通道发送数据Done")
}
func cachechan() {
	// 无缓冲通道 直接发送数据而无接收方时会出现panic
	//var c chan int
	c := make(chan int, 4)
	// 使用make而不是使用var创建，var创建一个nil 通道
	go recv(c)
	c <- 10
	c <- 11
	c <- 12
	fmt.Println("向通道发送数据Done")
}
func recv(c chan int) {
	val := <-c
	fmt.Printf("从通道接收到的数据为：%d \n", val)
}

/**
1、实现将 1 ~100数值发送到ch1
2、从ch1中取出数据乘以2 发送到ch2
3、打印ch2中数据
*/
func sendtochan() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * 2
		}
		close(ch2)

	}()
	for ret := range ch2 {
		fmt.Println(ret)
	}

}
