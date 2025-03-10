package cycle_array

import (
	"errors"
	"fmt"
)

type CycleArray struct {
	arr []int
	//环形队列 队头索引
	startIndex int
	//环形队列 队尾索引
	endIndex int
	// 容量大小
	size int
	// 元素个数
	count int
}

func NewCycleArray(size int) *CycleArray {
	return &CycleArray{
		arr:        make([]int, size),
		startIndex: 0,
		endIndex:   0,
		size:       size,
		count:      0,
	}
}

func (ca *CycleArray) isFull() bool {
	return ca.count == ca.size
}
func (ca *CycleArray) isEmpty() bool {
	return ca.count == 0
}
func (ca *CycleArray) resize(newSize int) {
	newArr := make([]int, newSize)
	for i := 0; i < ca.count; i++ {
		newArr[i] = ca.arr[(ca.startIndex+i)%ca.size]
	}
	ca.arr = newArr
	ca.startIndex = 0
	ca.endIndex = ca.count
	ca.size = newSize
}
func (ca *CycleArray) AddFirst(val int) {
	if ca.isFull() {
		ca.resize(ca.size * 2)
	}
	ca.startIndex = (ca.startIndex - 1 + ca.size) % ca.size
	ca.arr[ca.startIndex] = val
	ca.count++
}
func (ca *CycleArray) RemoveFirst() error {
	if ca.isEmpty() {
		return errors.New("cycle array is empty")
	}
	// 先去掉
	ca.arr[ca.startIndex] = 0
	// startIndex往右移动一位
	ca.startIndex = (ca.startIndex + 1) % ca.size
	ca.count--
	if ca.count > 0 && ca.count == ca.size/4 {
		ca.resize(ca.size / 2)
	}
	return nil
}
func (ca *CycleArray) RemoveLast() error {

	if ca.isEmpty() {
		return errors.New("cycle array is empty")
	}

	ca.endIndex = (ca.endIndex - 1 + ca.size) % ca.size
	ca.arr[ca.endIndex] = 0
	ca.count--
	if ca.count > 0 && ca.count == ca.size/4 {
		ca.resize(ca.size / 2)
	}
	return nil

}
func (ca *CycleArray) AddLast(val int) {
	if ca.isFull() {
		ca.resize(ca.size * 2)
	}
	ca.arr[ca.endIndex] = val
	ca.endIndex = (ca.endIndex + 1) % ca.size
	ca.count++
}
func (ca *CycleArray) Print() {
	for i := 0; i < ca.count; i++ {
		index := (ca.startIndex + i) % ca.size
		fmt.Printf("index is %d val is %d \r\n", index, ca.arr[index])
	}
}
