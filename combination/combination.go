package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// 组合模式

// Context 上下文
type Context struct {
}

// Component 组件接口
type Component interface {
	// 添加一个子组件
	Mount(c Component, components ...Component) error
	// 移除一个子组件
	Remove(c Component) error
	// 执行组件&子组件
	Do(ctx *Context) error
}

var _ Component = (*BaseComponent)(nil)

// BaseComponent 基础组件
// 实现 Add:添加一个子组件
// 实现 Remove:移除一个子组件
type BaseComponent struct {
	// 子组件列表
	ChildComponents []Component
}

// Mount 挂载一个子组件
func (bc *BaseComponent) Mount(c Component, components ...Component) (err error) {
	bc.ChildComponents = append(bc.ChildComponents, c)
	if len(components) == 0 {
		return
	}
	bc.ChildComponents = append(bc.ChildComponents, components...)
	return
}

// Remove 移除一个子组件
func (bc *BaseComponent) Remove(c Component) (err error) {
	if len(bc.ChildComponents) == 0 {
		return
	}
	for k, childComponent := range bc.ChildComponents {
		if c == childComponent {
			fmt.Println(runFuncName(), "移除:", reflect.TypeOf(childComponent))
			bc.ChildComponents = append(bc.ChildComponents[:k], bc.ChildComponents[k+1:]...)
		}
	}
	return
}

func (bc *BaseComponent) Do(ctx *Context) (err error) {
	// do nothing
	return
}

// ChildsDo 执行子组件
func (bc *BaseComponent) ChildsDo(ctx *Context) (err error) {
	// 执行子组件
	for _, childComponent := range bc.ChildComponents {
		if err = childComponent.Do(ctx); err != nil {
			return err
		}
	}
	return
}

// CHeckoutPageComponent 订单结算页面组件
type CheckoutPageComponent struct {
	// 合成复用基础组件
	BaseComponent
}

func (bc *CheckoutPageComponent) Do(ctx *Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(runFuncName(), "订单结算页面组件。。。")
	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写着
	return
}

// AddressComponent 地址组件
type AddressComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件 & 子组件
func (bc *AddressComponent) Do(ctx *Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(runFuncName(), "地址组件...")

	// 执行组件
	bc.ChildsDo(ctx)
	// 当前组件的业务逻辑写这
	return
}

// StoreComponent 店铺组件
type StoreComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *StoreComponent) Do(ctx *Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(runFuncName(), "店铺组件...")

	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写这

	return
}

// PayMethodComponent 支付方式组件
type PayMethodComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *PayMethodComponent) Do(ctx *Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(runFuncName(), "支付方式组件...")

	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写这

	return
}

// AftersaleComponent 售后组件
type AftersaleComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *AftersaleComponent) Do(ctx *Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(runFuncName(), "售后组件...")

	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写这

	return
}

// InvoiceComponent 发票组件
type InvoiceComponent struct {
	// 合成复用基础组件
	BaseComponent
}

// Do 执行组件&子组件
func (bc *InvoiceComponent) Do(ctx *Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(runFuncName(), "发票组件...")

	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写这

	return
}

func main() {
	// 初始化订单结算页面 这个大组件
	checkoutPage := CheckoutPageComponent{}

	// 挂载子组件
	storeComponent := &StoreComponent{}
	aftersale := &AftersaleComponent{}
	aftersale.Mount(
		&InvoiceComponent{},
	)

	storeComponent.Mount(
		&AddressComponent{},
		&PayMethodComponent{},
	)
	checkoutPage.Mount(
		storeComponent,
		aftersale,
	)

	// 移除组件测试
	//checkoutPage.Remove(&AddressComponent{})

	// 开始构建页面组件数据
	checkoutPage.Do(&Context{})
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
