package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 责任链模式

// Context Context
type Context struct {
}

// Handler 处理
type Handler interface {
	// 自身的业务
	Do(c *Context) error
}

// Next 抽象出来的 可被合成复用的结构体
// 使用链表结构体优化 -- 队列优化，先进先出
type Next struct {
	// 通过使用链表进行优化
	mu   sync.Mutex
	size int
	root *LinkNode // 首节点 -- 一直指向头节点
	tail *LinkNode // 尾结点 -- 一直指向尾结点
}

type LinkNode struct {
	next        *LinkNode
	nextHandler Handler
}

// SetNext 实现好的，可被复用的SetNext 方法
// 返回值是下一个对象 方便写成链式代码优雅
// 例如 nullHandler.SetNext(argumentsHandler).SetNext(signHandler).SetNext(frequentHandler)
func (n *Next) SetNext(h Handler) *Next {
	n.mu.Lock()
	defer n.mu.Unlock()

	// 创建新节点
	newNode := new(LinkNode)
	newNode.nextHandler = h

	// 判断是否为空
	if n.root == nil {
		// 创建节点
		n.root = newNode
		n.tail = n.root
	} else {
		// 存在节点
		// 添加节点 -- 让本身成为节点
		n.tail.next = newNode
		n.tail = newNode
	}
	n.size += 1
	return n
}

// Run 执行
func (n *Next) Run(c *Context) (err error) {
	// 由于go无继承的概念 这里无法执行当前handler的Do
	// n.Do(c)
	if n.root != nil {
		topNode := n.root
		for topNode.next != nil {
			if err = topNode.nextHandler.Do(c); err != nil {
				return
			}
			topNode = topNode.next
		}
	}
	return
}

// NullHandler 空Handler
// 由于go 无集成的概念 作为链式调用的第一个载体 设置实际的下一个对象
type NullHandler struct {
	// 合成复用Next 的`nextHandler` 成员属性、`SetNext`成员方法、`Run`成员方法
	Next
}

// Do 空Handler的Do
func (h *NullHandler) Do(c *Context) (err error) {
	// 空 Handler 这里什么也不做 只是载体 do nothing
	return
}

// ArgumentsHandler 校验参数的handler  -- 校验参数方法
type ArgumentsHandler struct {
	// 合成复用Next
	Next
}

// 校验参数逻辑
func (h *ArgumentsHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "校验参数成功。。。")
	return
}

// AddressInfoHandler 地址信息handler
type AddressInfoHandler struct {
	// 合成复用 Next
	Next
}

// Do 校验参数的逻辑
func (h *AddressInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取地址信息。。。")
	fmt.Println(runFuncName(), "地址信息校验。。。")
	return
}

// CartInfoHandler 获取购物车数据 handler
type CartInfoHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *CartInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取购物车数据。。。")
	return
}

//.....

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {
	nullHander := &NullHandler{}
	nullHander.SetNext(&ArgumentsHandler{}).
		SetNext(&AddressInfoHandler{}).
		SetNext(&CartInfoHandler{})
	// ... 无限扩展代码

	if err := nullHander.Run(&Context{}); err != nil {
		// 异常
		fmt.Println("Fail | Error:", err.Error())
		return
	}

	// 成功
	fmt.Println("Success")
	return
}
