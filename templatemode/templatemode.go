package main

import (
	"fmt"
	"runtime"
)

// 模板模式
const (
	// ConstActTypeTime 按时间抽奖类型
	ConstActTypeTime = 1 + iota
	// ConstActTypeTimes 按抽奖次数抽奖
	ConstActTypeTimes
	// ConstActTypeAmount 按数额范围区间抽奖
	ConstActTypeAmount
)

// Context 上下文
type Context struct {
	ActInfo *ActInfo
}

type ActInfo struct {
	// 活动抽奖类型 1：按时间抽奖 2：按抽奖次数抽奖 3：按数额范围区间抽奖
	ActivityType int32
	// 其他字段略
}

// BehaviorInterface 不同抽奖类型的行为
type BehaviorInterface interface {
	// 其他参数校验（不同活动类型实现不同）
	checkParams(ctx *Context) error
	// 获取 node 奖品信息（不同活动类型实现不同）
	getPrizesByNode(ctx *Context) error
}

var _ BehaviorInterface = (*TimeDraw)(nil)

// TimeDraw 具体抽奖行为
// 按时间抽奖类型 比如红包雨
type TimeDraw struct {
}

// checkParams 其他参数校验(不同活动类型实现不同)
func (draw TimeDraw) checkParams(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "do nothing(抽取该场次的奖品即可，无需其他逻辑)")
	return
}

// getPrizesBYNode 获取node 奖品信息(不同活动类型实现不同)
func (draw TimeDraw) getPrizesByNode(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "1.判断是该用户第几次抽奖。。。")
	fmt.Println(runFuncName(), "2.获取对应node的奖品信息。。。")
	fmt.Println(runFuncName(), "3.复写所有奖品信息(抽取该node节点的奖品)。。。")
	return
}

var _ BehaviorInterface = (*AmountDraw)(nil)

// 具体抽奖行为
// 按数额范围区间抽奖，比如订单金额刮奖
type AmountDraw struct {
}

// checkParams 其他参数校验(不同活动类型实现不同)
func (draw *AmountDraw) checkParams(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "按数额范围区间抽奖：特殊参数校验。。。")
	return
}

func (draw *AmountDraw) getPrizesByNode(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "1. 判断属于哪个数额区间...")
	fmt.Println(runFuncName(), "2. 获取对应node的奖品信息...")
	fmt.Println(runFuncName(), "3. 复写原所有奖品信息(抽取该node节点的奖品)...")
	return
}

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// Lottery 抽奖模板
type Lottery struct {
	concreteBehavior BehaviorInterface
}

// Run 抽奖算法
// 稳定不变的算法步骤
func (lottery *Lottery) Run(ctx *Context) (err error) {
	// 具体方法：校验活动编号 (serial_no) 是否存在，并获取活动信息
	if err = lottery.checkSerialNo(ctx); err != nil {
		return
	}
	// ....
	return
}

func (lottery *Lottery) checkSerialNo(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "校验活动编号(serial_no)是否存在、并获取活动信息表")
	// 获取活动信息伪代码
	ctx.ActInfo = &ActInfo{
		// 假设当前的活动类型为按抽奖次数抽奖
		ActivityType: ConstActTypeTime,
	}
	// 具体行为
	switch ctx.ActInfo.ActivityType {
	case ConstActTypeTime:
		// 按时间抽象
		lottery.concreteBehavior = &TimeDraw{}
	case ConstActTypeTimes:
		// 按数额范围区间抽奖
		lottery.concreteBehavior = &AmountDraw{}
	}
	return
}

func main() {
	(&Lottery{}).Run(&Context{})
}
