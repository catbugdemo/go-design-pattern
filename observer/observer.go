package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// 观察者模式

// Observable 观察者
type Observable interface {
	Attach(observer ...ObserverInterface) Observable
	Detach(observer ObserverInterface) Observable
	Notify() error
}

// ObserverInterface 定义一个被观察者接口
type ObserverInterface interface {
	Do(o Observable) error
}

var _ Observable = (*ObservableConcrete)(nil)

// ObservableConcrete 一个具体的订单状态变化的被观察者
type ObservableConcrete struct {
	observerList []ObserverInterface
}

// Attach 注册观察者
func (o *ObservableConcrete) Attach(observer ...ObserverInterface) Observable {
	o.observerList = append(o.observerList, observer...)
	return o
}

func (o *ObservableConcrete) Detach(observer ObserverInterface) Observable {
	if len(o.observerList) == 0 {
		return o
	}
	for k, observerItem := range o.observerList {
		if observer == observerItem {
			fmt.Println(runFuncName(), "注销：", reflect.TypeOf(observer))
			o.observerList = append(o.observerList[:k], o.observerList[k+1:]...)
		}
	}
	return o
}

// 通知观察者
func (o *ObservableConcrete) Notify() error {
	// code ...
	for _, observer := range o.observerList {
		if err := observer.Do(o); err != nil {
			return err
		}
	}
	return nil
}

// OrderStatus 修改订单状态
type OrderStatus struct {
}

// Do 具体业务
func (order *OrderStatus) Do(o Observable) error {
	// code ...
	fmt.Println(runFuncName(), "修改订单状态...")
	return nil
}

// OrderStatusLog 记录订单状态变更日志
type OrderStatusLog struct {
}

// Do 具体业务
func (order *OrderStatusLog) Do(o Observable) error {
	fmt.Println(runFuncName(), "退优惠券")
	return nil
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {
	watch := &ObservableConcrete{}

	watch.Attach(&OrderStatus{}).
		Attach(&OrderStatusLog{})

	watch.Notify()
}
