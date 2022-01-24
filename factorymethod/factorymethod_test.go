package factorymethod

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	facotry := PlusOperatorFactory{}
	if compute(facotry, 1, 2) != 3 {
		t.Fatal("error with factory methdo pattern")
	}
}
