package linklist

import "fmt"

func IsUglyNumber(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n = n / 2
	}
	for n%3 == 0 {
		n = n / 3
	}
	for n%5 == 0 {
		n = n / 5
	}
	return n == 1
}
func UglyNumberOfN(n int) int {
	// 2 ,3 ,5倍数的链表
	product2, product3, product5 := 1, 1, 1
	// 链表头
	p2, p3, p5 := 1, 1, 1

	uglys := make([]int, n+1)
	p := 1
	for p <= n {
		min := min(product2, product3, product5)
		uglys[p] = min

		if min == product2 {
			product2 = 2 * uglys[p2]
			p2++
		}
		if min == product3 {
			product3 = 3 * uglys[p3]
			p3++
		}
		if min == product5 {
			product5 = 5 * uglys[p5]
			p5++
		}

		p++
	}
	fmt.Printf("uglys is %+v \r\n", uglys)
	return uglys[n]
}
