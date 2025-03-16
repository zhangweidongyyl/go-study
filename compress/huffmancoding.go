package compress

import (
	"container/heap"
)

// HuffmanNode 哈夫曼树节点
type HuffmanNode struct {
	value rune // 数字
	freq  int  // 频率
	left  *HuffmanNode
	right *HuffmanNode
}

// PriorityQueue 优先队列（最小堆）
type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].freq < pq[j].freq }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// buildHuffmanTree 创建huffman tree 按字符出现频率的最小堆进行 创建huffman tree ,频率最小的在root
// 字符都出现在叶子节点，父节点的freq为子节点的freq之和
func buildHuffmanTree(data string) *HuffmanNode {
	// 统计频率
	freqMap := make(map[rune]int)
	for _, c := range data {
		freqMap[c]++
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	// 将所有值及其频率加入优先队列
	for value, freq := range freqMap {
		heap.Push(pq, &HuffmanNode{value: value, freq: freq})
	}

	// 构建哈夫曼树
	for pq.Len() > 1 {
		left := heap.Pop(pq).(*HuffmanNode)
		right := heap.Pop(pq).(*HuffmanNode)
		parent := &HuffmanNode{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}
		heap.Push(pq, parent)
	}

	return heap.Pop(pq).(*HuffmanNode)
}

// 生成哈夫曼编码表
func buildHuffmanCodes(root *HuffmanNode, code string, codes map[rune]string) {
	if root == nil {
		return
	}

	// 叶子节点，存储编码
	if root.left == nil && root.right == nil {
		codes[root.value] = code
		return
	}

	// 递归生成编码
	buildHuffmanCodes(root.left, code+"0", codes)
	buildHuffmanCodes(root.right, code+"1", codes)
}

// HuffmanCompress 压缩函数
func HuffmanCompress(data string, codes map[rune]string) []byte {
	var compressed []byte
	var buffer byte
	var count uint

	for _, char := range data {
		code := codes[char]
		for _, bit := range code {
			if bit == '1' {
				buffer |= 1 << (7 - count)
			}
			count++
			if count == 8 {
				compressed = append(compressed, buffer)
				buffer = 0
				count = 0
			}
		}
	}

	// 处理剩余的比特
	if count > 0 {
		compressed = append(compressed, buffer)
	}

	return compressed
}

// HuffmanDecompress 哈夫曼编码解压
// 解压函数
func HuffmanDecompress(compressed []byte, root *HuffmanNode, length int) []rune {
	// 将字节流转换为二进制字符串
	var bitStream string
	for _, b := range compressed {
		for i := 7; i >= 0; i-- {
			if b&(1<<i) != 0 {
				bitStream += "1"
			} else {
				bitStream += "0"
			}
		}
	}

	// 解压
	var result []rune
	currentNode := root
	for _, bit := range bitStream {
		if bit == '0' {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}

		// 到达叶子节点
		if currentNode.left == nil && currentNode.right == nil {
			result = append(result, currentNode.value)
			currentNode = root // 重置为根节点
			if len(result) == length {
				break
			}
		}
	}

	return result
}
