# hqc-cache
基于主从架构的分布式缓存系统
要实现的功能：

1.数据一致性(最终一致性,进程通信(采用tcp,自定义协议))

2.HA(心跳,分布式选举)

3.容灾(数据持久化)

4.读写分离(主节点写 ，从节点读)。

5.客户端(增删改查接口，注册，心跳,负载均衡，先写java版本)

6.支持常见数据结构(map list array 等)
