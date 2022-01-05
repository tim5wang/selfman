
[NSQ](https://github.com/nsqio/nsq): 是一个go语言实现的消息队列

文档地址：https://nsq.io/

# [Design](https://nsq.io/overview/design.html)

## 特性

支持拓扑，使高可用性和消除(SOPFs)单点故障

满足更强大的信息交付保证需求

绑定单个过程的内存占用面（通过将某些消息持久化到磁盘）

大大简化了生产者和消费者的配置要求

提供直接的升级路径

提高效率

## 消息模型

kafka 会把topic分成多个partition, 也会把Client分成多个Group，每个Group会分担所有partition的消息，也就是有多少Group就会把消息消费多少遍。

kafka会把消息直接写到硬盘，使用了

# 参考
https://zhuanlan.zhihu.com/p/46421050
