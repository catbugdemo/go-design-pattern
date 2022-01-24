package factorymethod

// 工厂方法其实就只是两层 interface 封装

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a int
	b int
}

//var _ Operator = (*OperatorBase)(nil)

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 工厂方法类
type PlusOperatorFactory struct {
}

var _ OperatorFactory = (*PlusOperatorFactory)(nil)

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (o *PlusOperator) Result() int {
	return o.a + o.b
}
