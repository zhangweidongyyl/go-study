package main

import (
	"study.com/study/compress"
)

func main() {
	datas := []uint32{10, 11, 12, 12, 12, 15, 16, 16, 18, 15, 11, 11, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	fileName := "test.md"
	outputFileName := "test1.md"
	compress.CreateFile(fileName, datas)
	compress.StreamReadEncodeNoMemory(fileName, outputFileName)
	compress.StreamReadDecodeNoMemory(outputFileName)
	//encodedResult := compress.RuneLengthEncode(datas)
	//b := compress.EncodeToBinary(encodedResult)
	//bres := compress.DecodeBinaryToDomains(b)
	//fmt.Printf("bres is %+v \r\n", bres)
	//fmt.Printf("datas length is %d \r\n", unsafe.Sizeof(b))
	//fmt.Printf("encodeResult is %+v and memory is %d \r\n", encodedResult, unsafe.Sizeof(encodedResult))
	//res := compress.RuneLengthDecode(encodedResult)
	//fmt.Printf("res is %+v \r\n", res)
	//// 示例数据
	//data := []int{10, 11, 12, 12, 12, 15, 16, 16, 18}
	//
	//// 压缩
	//compressed := linklist.Compress(data)
	//fmt.Println("压缩后的数据:", compressed)
	//
	//// 解压
	//decompressed := linklist.Decompress(compressed)
	//fmt.Println("解压后的数据:", decompressed)
	//
	//// 检查解压是否正确
	//if fmt.Sprint(data) == fmt.Sprint(decompressed) {
	//	fmt.Println("解压成功！")
	//} else {
	//	fmt.Println("解压失败！")
	//}

	//numArray := linklist.NewNumArray([]int{1, 2, 3, 4, 5, 6})
	//numArray.Increment(1, 3, 2)
	//res1 := numArray.GetOriginNumsByDiffs()
	//fmt.Printf("numArray is %+v and sunRangeRes is %d \r\n", numArray, res1)
	//os.Exit(1)
	//gochan.TestChannel()
	//data := make(map[string]interface{})
	//data["key"] = 200
	//os.Exit(1)
	//grammar.GenerateTrace()
	//os.Exit(1)
	//ca := cycle_array.NewCycleArray(10)
	//ca.AddFirst(2)
	//ca.AddFirst(22)
	//ca.AddLast(887)
	//ca.RemoveLast()
	//ca.Print()
	//nums := []int{2, 1, 2, 4, 3}
	//res := stack.CalculateGreaterElement(nums)
	//fmt.Printf("CalculateGreaterElement res is %+v \r\n", res)
	//os.Exit(1)
	//arrayQueue := myqueue.NewDefaultArrayQueue()
	//arrayQueue.Enqueue(1)
	//arrayQueue.Enqueue(2)
	//arrayQueue.Enqueue(3)
	//arrayQueue.Enqueue(4)
	//arrayQueue.Dequeue()
	//arrayQueue.Enqueue(5)
	//recentNumber := myqueue2.NewRecentNumber()
	//recentNumber.Ping(1)
	//recentNumber.Ping(4)
	//recentNumber.Ping(2000)
	//recentNumberRes := recentNumber.Ping(5000)
	//fmt.Printf("recentNumberRes is %+v \r\n", recentNumberRes)
	//os.Exit(1)
	//myqueue := myqueue2.MyQueue{}
	//myqueue.Push(1)
	//myqueue.Push(2)
	//myqueue.Push(3)
	//fmt.Println(myqueue.Pop())
	//os.Exit(1)
	//myStack := stack.NewMyStack()
	//myStack.Push(1)
	//myStack.Push(2)
	//fmt.Println(myStack.Pop())
	//myStack1 := stack.NewMyStack1()
	//myStack1.Push(1)
	//myStack1.Push(2)
	//fmt.Println(myStack1.Pop())
	//
	//os.Exit(1)
	//stk := list.New()
	//stk.PushFront(1)
	//stk.PushFront(2)
	//fmt.Println(stk.Front().Value)
	//os.Exit(1)
	//
	//head := make([]int, 0)
	//head = append(head, 1)
	//head = append(head, 2)
	//for _, ele := range head {
	//	fmt.Printf("ele is %d \r\n", ele)
	//}
	//fmt.Printf("peek is %d \r\n", head[0])
	//os.Exit(1)
	//mystack := stack.NewMyStack()
	//mystack.Push(1)
	//mystack.Push(2)
	//fmt.Println(mystack.Peek())
	//os.Exit(1)
	//rpnRes := stack.InfixToRpn("3 + 5 * 4 + 8 / 2")
	//evalRes := stack.EvalRpn(rpnRes)
	//fmt.Printf("rpnRes is %+v \r\n", rpnRes)
	//fmt.Printf("evalRes is %+v \r\n", evalRes)
	//os.Exit(1)
	//// 交换后 1 5 2 4 3
	//l1 := linklist.CreateLinkNode([]int{1, 2, 3, 4, 5})
	//l1res := stack.ReOrderLinkNode(l1)
	//
	//linklist.PrintLinkNode(l1res)
	//os.Exit(1)
	//unixPath := "/./../.././//users/code/in/out/"
	//simplifyPath := stack.SimplifyPath(unixPath)
	//fmt.Printf("simplifyPath is %s \r\n", simplifyPath)
	//os.Exit(1)
	////trips = [[2,1,5],[3,3,7]], capacity = 4
	//trips := make([][]int, 0)
	//trip1 := []int{2, 1, 5}
	//trip2 := []int{3, 3, 7}
	//trips = append(trips, trip1, trip2)
	//carPoolingRes := linklist.CarPooling(trips, 6)
	//fmt.Printf("carPoolingRes is %+v \r\n", carPoolingRes)
	//os.Exit(1)
	//// 1-2 航班订了10个座位
	//booking1 := []int{1, 2, 10}
	//
	//booking2 := []int{2, 3, 20}
	//booking3 := []int{3, 3, 10}
	////   预订记录    航班  座位数
	//// booking1    1，2   10
	//// booking2    2，3   20
	//// booking3    3，3   10
	//// 1 10
	//// 2 30
	//// 3 30
	//bookings := make([][]int, 0)
	//bookings = append(bookings, booking1, booking2, booking3)
	//calcRes := linklist.CalcFlightBookings(bookings, 3)
	//fmt.Printf("calcRes is %+v \r\n", calcRes)
	//os.Exit(1)
	//// 1 2 3 4 5 6
	//// 1 4 5 6 5 6
	//// 1 3 1 -1 1
	//numArray := linklist.NewNumArray([]int{1, 2, 3, 4, 5, 6})
	//numArray.Increment(1, 3, 2)
	//res1 := numArray.GetOriginNumsByDiffs()
	//fmt.Printf("numArray is %+v and sunRangeRes is %d \r\n", numArray, res1)
	//os.Exit(1)
	//piles := []int{3, 6, 7, 11}
	//pilesRes := linklist.MinEatSpeed(piles, 8)
	//fmt.Printf("pilesRes is %+v \r\n", pilesRes)
	//os.Exit(1)
	//binarySearchNums := []int{5, 6, 8, 8, 10}
	//binarySearchRes := linklist.RightBound(binarySearchNums, 8)
	//fmt.Printf("binarySearchRes is %+v \r\n", binarySearchRes)
	//os.Exit(1)
	//pNums := []int{10, 5, 2, 6}
	//k := 100
	//pNUmsRes := linklist.FindProductLessThanK(pNums, k)
	//fmt.Printf("pNumsRes is %+v \r\n", pNUmsRes)
	//
	//os.Exit(1)
	//
	//nums := []int{1, 2, 2, 2, 3}
	//minOperationsRes := linklist.FindMinOperations(nums, 6)
	//fmt.Printf("min operations is %+v \r\n", minOperationsRes)
	//os.Exit(1)
	//findLongestSubString := linklist.FindLongestSubString("abcbacabcdef")
	//fmt.Printf("findLongestSubString is %+v \r\n", findLongestSubString)
	//os.Exit(1)
	//findStringPosRes := linklist.FindStringPosition("sbbssb", "bs")
	//fmt.Printf("findStringPosRes is %+v \r\n", findStringPosRes)
	//os.Exit(1)
	//checkContainsString := linklist.CheckContainsString("koisb", "bs")
	//fmt.Printf("checkContainsString is %+v \r\n", checkContainsString)
	//os.Exit(1)
	//minStringRes := linklist.MinString("45abcccsd", "bs")
	//fmt.Printf("minStringRes is %+v \r\n", minStringRes)
	//os.Exit(1)
	////  1  2  3
	////  8  9  4
	////  7  6  5
	//generateMatrixRes := linklist.GenerateMatrix(3)
	//fmt.Printf("generateMatrixRes is %+v \r\n", generateMatrixRes)
	//os.Exit(1)
	////   1 2 4 5
	////   2 3 5 7
	////   9 8 7 6
	////   7 5 4 3
	//// 翻转后
	////   1 2 9 7
	////   2 3 8 5
	//
	//row1 := []int{1, 2, 4, 5}
	//row2 := []int{2, 3, 5, 7}
	//row3 := []int{9, 8, 7, 6}
	//row4 := []int{7, 5, 4, 3}
	//matrix := [][]int{
	//	row1, row2, row3, row4,
	//}
	//rotateRes := linklist.RotateMatrix(matrix)
	//fmt.Printf("rotateRes is %+v \r\n", rotateRes)
	//
	//os.Exit(1)
	//
	//str1 := "12abccba2323d123321"
	//r := linklist.LongestPalidromeString(str1)
	//fmt.Printf("r is %s \r\n", r)
	//os.Exit(1)
	//str := linklist.ReverseString([]rune("acdsd"))
	//fmt.Printf("str is %s \r\n", str)
	//os.Exit(0)
	//res := linklist.TwoSum([]int{2, 7, 11, 19}, 9)
	//fmt.Printf("res is %+v \r\n", res)
	// linklist相关
	//1、merge two link list
	//
	//l1 := linklist.CreateLinkNode([]int{1, 3, 5})
	//l2 := linklist.CreateLinkNode([]int{2, 6, 9, 14, 10})
	//l3 := linklist.MergeTwoList(l1, l2)
	////linklist.PrintLinkNode(l3)
	//
	//// 2 、 partition link list
	//l4 := linklist.PartitionList(l3, 5)
	//linklist.PrintLinkNode(l4)

	//l1 := linklist.CreateLinkNode([]int{1, 3, 5})
	//l2 := linklist.CreateLinkNode([]int{2, 6, 9, 14, 8})
	//
	//res := linklist.MiddleNode(l2)
	//linklist.PrintLinkNode(res)
	//
	//l2Repeat := linklist.CreateLinkNode([]int{2, 2, 3, 6, 6, 8})
	//
	//resUniq := linklist.DeleteRepeat2(l2Repeat)
	//linklist.PrintLinkNode(resUniq)
	//res := linklist.RemoveFromEnd(l2, 2)
	//linklist.PrintLinkNode(res)

	//l1 := linklist.CreateLinkNode([]int{1, 3, 5})
	//l2 := linklist.CreateLinkNode([]int{2, 6, 9, 14})
	//list := make([]*linklist.LinkNode, 0)
	//list = append(list, l1, l2)
	//merge k list 使用优先级队列 合并k个升序队列
	//res := linklist.MergeKList(list)
	////
	//linklist.PrintLinkNode(res)
	//
	////fmt.Printf("ugly number %d is %d \r\n", 14, linklist.UglyNumberOfN(14))
	//fmt.Printf("res is %+v \r\n", linklist.IsUglyNumber(6))

	//matrix := make([][]int, 0)
	//e1 := []int{1, 2, 3, 4}
	//e2 := []int{4, 5, 6, 7}
	//matrix = append(matrix, e1, e2)
	//k := 7
	//
	//res := linklist.FindKThInMatrix(matrix, k)
	//fmt.Printf("the %d min number in matrix is %d", k, res)

	//e1 := []int{1, 2, 7}
	//e2 := []int{2, 4, 6}
	//k := 3
	//
	//res := linklist.FindKThSumInMatrix(e1, e2, k)
	//fmt.Printf("the %d min number in matrix is %d", k, res)

	//e1 := []int{1, 2, 7, 6, 9, 10}
	//l1 := linklist.CreateLinkNode(e1)
	//
	////res := linklist.ReverseN(l1, 3)
	////res := linklist.ReverseBetween(l1, 2, 5)
	//linklist.Traverse(l1)
	//
	//
	//
	//checkIsPalidromeString := linklist.CheckIsPalidromeString("abccbafdg")
	//fmt.Printf("check is palidrome string is %+v \r\n", checkIsPalidromeString)

	//l1 := linklist.CreateLinkNode([]int{1, 2, 3, 2, 1})
	//checkIsPalidromeLink := linklist.CheckIsPalidromeLink1(l1)
	//fmt.Printf("check is palidrome link is %+v \r\n", checkIsPalidromeLink)
	//nums := []int{0, 0, 2, 2, 3, 4, 5, 6, 6, 7, 8, 8, 8, 9, 11}
	//res := linklist.MoveZero(nums)
	//fmt.Printf("MoveZero result is %+v \r\n", res)
	//l1 := linklist.CreateLinkNode(nums)
	//res1 := linklist.DeleteRepeatInPlaceForLink(l1)
	//linklist.PrintLinkNode(res1)
	//res := linklist.DeleteRepeatInPlaceForArray(nums)
	//fmt.Printf("res is %+v \r\n", res)
	//res := linklist.ReverseKGroup(l1, 2)
	//linklist.PrintLinkNode(res)
	//e2 := []int{2, 4, 6}
	//l2 := linklist.CreateLinkNode(e2)
	//res := linklist.AddTwoNumbers2(l1, l2)
	//linklist.PrintLinkNode(res)

	//treenode.PrintPq()
	//ntreeroot := dfs.CreateNTree()
	//dfs.TraverseNTree(ntreeroot)

	//ntreeroot := bfs.CreateNTree()
	//bfs.NtreeLevelOrderTraverse3(ntreeroot)

	//root := dfs.CreateMinDepthTree()
	//res := dfs.FindAllPath(root)
	//for index, path := range res {
	//	fmt.Printf("path index is %d and path is %+v \r\n", index, path)
	//}

	//root := bfs.CreateMinDepthTree()
	//res := bfs.LevelOrderFindAllPath(root)
	//for index, path := range res {
	//	fmt.Printf("path index is %d and path is %+v \r\n", index, path)
	//}
	//minDepth := dfs.MinDepth(root)
	//fmt.Println(minDepth)
	//root := dfs.CreateTree()
	//root.Traverse()
	//root := bfs.CreateTree()
	//bfs.LevelOrderTraverse1(root)
	//bfs.LevelOrderTraverse2(root)
	//bfs.LevelOrderTraverse3(root)

	//root := bfs.CreateMinDepthTree()
	//fmt.Println(bfs.LevelOrderMinDepth(root))

}
