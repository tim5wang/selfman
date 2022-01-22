（Log-Structured-Merge-Tree 日志结构合并树）


LSM 树是一种非常强大的数据结构，现今已经被广泛应用到工业界的产品中了，如：
HBase、Cassandra、LevelDB、RocksDB、MongoDB、Wired Tiger

https://juejin.cn/post/6844903863758094343


一个重要的前置知识是:
磁盘的顺序写的吞吐量甚至能够超过内存的随机写的吞吐量，LSM树正是利用了这一点，它通过讲磁盘随机写操作转化为顺序写操作，从而将随机写操作的吞吐量提高了几个数量级


LSM 的问题主要在于牺牲了一定读性能，其中 CheckExist的hit结果是False Positive时的代价是很大的，以为着扫库，因此一般采用布隆过滤器先做筛选。

[kafka数据是持久化原理](https://www.jianshu.com/p/1bdc181a7ebd)：
> 写数据：过内存S大小（S可设置）的数据后，直接刷进磁盘，追加写入文件；
读数据：根据offset读取位置之后S大小的数据，进内存；
删数据：直接删磁盘文件（segment file），先删老文件（可设置）。

[raft协议](https://juejin.cn/post/6907151199141625870#heading-32)

