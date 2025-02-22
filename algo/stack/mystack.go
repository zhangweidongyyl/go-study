package stack

import (
	"strconv"
	"strings"
	"study.com/study/algo/linklist"
)

type MyStack struct {
	elements []int
}

func NewMyStack() MyStack {
	return MyStack{elements: make([]int, 0)}
}
func (this *MyStack) Push(ele int) {
	this.elements = append(this.elements, ele)
}
func (this *MyStack) Pop() int {
	if len(this.elements) > 0 {
		ele := this.elements[len(this.elements)-1]
		this.elements = this.elements[:len(this.elements)-1]
		return ele
	}
	return -1
}
func (this *MyStack) Peek() int {
	if len(this.elements) > 0 {
		ele := this.elements[len(this.elements)-1]
		return ele
	}
	return -1
}
func (this *MyStack) Empty() bool {
	return len(this.elements) == 0
}

// InfixToRpn 逆波兰算法 中缀表达式 转换为 逆波兰表达式
func InfixToRpn(str string) string {
	// 具体输出的
	output := make([]string, 0)
	// 左右括号和操作符这类的字符
	operations := make([]string, 0)

	tokens := strings.Fields(str)
	for _, token := range tokens {
		// 遇到左括号
		if token == "(" {
			// 左括号
			operations = append(operations, token)
		} else if token == ")" {
			// 右括号
			for len(operations) > 0 && operations[len(operations)-1] != "(" {
				output = append(output, operations[len(operations)-1])
				operations = operations[:len(operations)-1]
			}
			operations = operations[:len(operations)-1] // 弹出左括号
		} else if !isOperation(token) {
			// 非操作符 具体是数字
			output = append(output, token)
		} else {
			// 具体的操作符 + - * / 这些
			// 栈里的 符号优先级比当前符号 的优先级高的时候 提前将 栈的 操作符压入 output

			for len(operations) > 0 && getOperationCharPriority(operations[len(operations)-1]) >= getOperationCharPriority(token) {
				output = append(output, operations[len(operations)-1])
				operations = operations[:len(operations)-1]
			}
			operations = append(operations, token)
		}
	}

	for len(operations) > 0 {
		output = append(output, operations[len(operations)-1])
		operations = operations[:len(operations)-1]
	}

	return strings.Join(output, " ")
}

func isOperation(c string) bool {
	if c == "+" || c == "-" || c == "*" || c == "/" {
		return true
	}
	return false
}
func getOperationCharPriority(c string) int {
	switch c {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func EvalRpn(rpn string) int {
	stack := make([]int, 0)
	tokens := strings.Fields(rpn)
	for _, token := range tokens {
		if strings.Contains("+-*/", token) {
			// 此时是操作符了
			a, b := stack[len(stack)-1], stack[len(stack)-2]

			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, b+a)
			case "-":
				stack = append(stack, b-a)
			case "*":
				stack = append(stack, b*a)
			case "/":
				stack = append(stack, b/a)

			}
		} else {
			// 此时是数字
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

// IsBracketValid 判断括号是否成对出现 并符合括号的规范
func IsBracketValid(str string) bool {
	stack := make([]rune, 0)
	for _, s := range str {
		// 左括号全部压入栈中
		if s == '(' || s == '[' || s == '{' {
			stack = append(stack, s)
		} else {
			// 遇到右括号
			// s 此时为右括号，当前栈顶为对应的左括号时 则pop弹出
			if len(stack) > 0 && getLeftBracketByRight(s) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	// 看是否所有左括号都被匹配了
	return len(stack) == 0
}

func getLeftBracketByRight(c rune) rune {
	if c == ')' {
		return '('
	} else if c == ']' {
		return '['
	} else if c == '}' {
		return '{'
	}
	return '0'
}

// SimplifyPath 简化给定的 unix文件路径
func SimplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stk := make([]string, 0)
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			// 如果 path里面有 .. 路径  则将其弹出
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}

			continue
		}

		stk = append(stk, part)
	}
	// 栈中存储的文件夹组成路径
	res := ""
	for _, dir := range stk {
		res += "/" + dir
	}
	if res == "" {
		return "/"
	}
	return res
}

// ReOrderLinkNode 对单链表进行重新排序
// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
func ReOrderLinkNode(head *linklist.LinkNode) *linklist.LinkNode {

	// 利用栈的 先进后出特性 将节点全部压进去，用于后续调整顺序
	stack := make([]*linklist.LinkNode, 0)

	p := head
	for p != nil {
		stack = append(stack, p)
		p = p.Next

	}

	p = head
	for p != nil {
		// 取出最后一个节点
		lastNode := stack[len(stack)-1]
		// 更新栈
		stack = stack[:len(stack)-1]

		nextNode := p.Next
		if lastNode == nextNode || lastNode.Next == nextNode {
			// 此时交换到最后了  需要考虑 奇数偶数个数的链表
			lastNode.Next = nil
			break
		}
		// 重新赋值
		p.Next = lastNode
		lastNode.Next = nextNode

		p = nextNode
	}

	return head
}
