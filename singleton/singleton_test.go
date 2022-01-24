package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

const parCount = 100

// 协程通道
func TestTestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]Singleton{}
	// 相当于生产者
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	// 关闭 channel 所有协程同时开始运行
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
