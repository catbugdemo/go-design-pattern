package main

// 观察者模式

// Observable 被观察者
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
}

func (o *ObservableConcrete) Notify() error {
	//TODO implement me
	panic("implement me")
}

// 胡
