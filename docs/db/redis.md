# Everything about Redis

**redis**是一个基于内存、支持持久化、事务、内置有多种数据结构的存储系统，可以作为数据库、消息系统、缓存系统。

QPS=10w/s，TPS=8w/s

## 数据结构

### string

### hash

### list

秒杀

### set

### sortedset

zset，元素有顺序且不能重复，有2种实现方式：

- ziplist，
- skiplist

### bloom filter

### bitmap

### hyperloglog

### geo

## 基础操作

### 指令

- `keys` vs. `scan`

### 配置

- 身份验证
- 白名单
- 禁用恶意指令集
-

## 核心机制

### 事务

- redis 事务不支持回滚
- redis 执行事务时不会被其他客户端打断

### 管道

### 持久化

RDB 和 AOF，

### 主从同步

## 应用场景

### 分布式锁

1. setnx+expire

### 缓存系统

首页数据，热点数据都需要做缓存，缓存触发 定时任务定时

### 排行系统

热帖、榜单、

### 限时系统


### 计数系统

院子递增`incrby`，秒杀、分布式序列号、调用限制（发送短信、接口调用）

### 消息队列
