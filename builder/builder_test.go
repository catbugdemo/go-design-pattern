package builder

import "testing"

// 可以这么说，创建者模型就是基础结构体类型进行 interface ，然后对其方法进行分类同时也进行 interface

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123" {
		t.Fatalf("Builder1 fail expect 123 acture %s", res)
	}
}
