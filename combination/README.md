
## 什么是「组合模式」？

> 一个具有层级关系的对象由一系列拥有父子关系的对象通过树形结构组成

组合模式的优势:

- 所见即所码：你所看见的代码结构就是业务真实的层级关系，比如 ui界面你真实看到那样
- 高度封装：单一职责
- 可复用：不同业务场景，相同的组件可被重复

## 什么真实业务场景可以用「组合模式」？

满足如下要求的所有场景：
> Get 请求获取页面数据的所有接口

前端大行组件化的当今，我们再写后端接口代码的时候还是按照业务思路一头写到尾吗？ 我们是否可以思索「后端接口业务代码如何可以简单快速组件化？」，答案是肯定的， 这就是「组合模式」的作用。

我们利用组合模式的定义和前端模块的划分去构建后端业务代码结构：

- 前端单个模块 -> 对应后端:具体单个类-> 封装的过程
- 前端模块父子组件 -> 对应后端：父类内部持有多个子类(非继承关系，合成复用关系)->父子关系的树形结构体

> 我们有哪些真实业务场景可以用「组合模式」？

从页面的展示形式上，可以看出：

- 页面由多个模块构成，比如：
    - 地址模块
    - 支付方式模块
    - 店铺模块
    - 发票模块
    - 优惠券模块
    - 某豆模块
    - 礼品卡模块
    - 订单详细金额模块
- 单个模块可由多个子模块构成
    - 店铺模块又由如下模块构成
        - 商品模块
        - 售后模块
        - 优惠模块
        - 物流模块

## 怎么用「组合模式」？

## 代码建模

责任链模式主要类主要包含如下特性：

- 成员属性
    - `ChildComponents`：子组件列表 -> 稳定不变的
- 成员方法
    - `Mount`:添加一个子组件 -> 稳定不变的
    - `Remove`:移除一个子组件 -> 稳定不变的
    - `Do`:执行组件&子组件 -> 变化

## 结语

最后总结下，「组合模式」抽象过程的核心是：

- 按模块划分：业务逻辑归归类，收敛的过程
- 父子关系(树)：把收敛之后的业务对象按父子关系绑定，依次被执行。

与「责任链模式」的区别

- 责任链模式：链表
- 组合模式：树
