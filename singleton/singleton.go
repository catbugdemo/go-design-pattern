package singleton

import "sync"

// 单例模式
// 1.某一个类只能有一个实例 2.它必须自行创建这个实例 3.它必须自行向整个系统提供这一实例

type Singleton interface {
	foo()
}

type singleton struct {
}

// 继承合理性验证
var _ Singleton = (*singleton)(nil)

func (s *singleton) foo() {
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
