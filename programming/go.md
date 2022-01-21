# 编程语言

- 变量
  - 常量
  - 数据类型
- 语句
  - 操作符
  - 表达式
    - 数值计算
    - 布尔计算
    - 字符拼接
  - 关键字
  - 声明语句
    - 变量声明
    - 函数声明
    - 类型声明
  - 赋值语句
  - 控制语句
    - 判断语句
    - 循环语句
    - 选择语句
- 函数
- 类型
  - 属性
  - 方法
- 集合
  - 可变数组
  - 字符串
  - 散列表
- 并发
- 特性
  - 面向函数
  - 面向对象
  - 单元测试
    - 逻辑测试函数
    - 基准测试函数
    - 示例函数
  - 错误处理
  - 代码管理

## 概念理解

### 变量

变量源于数学，用于短暂的保存数值，

### 语句

## 核心原理

## 标准库

### Json

### Yaml

---

# 编程规范

- 命名(naming)
  - 避免使用内置名称
  - 避免使用信息量不足的名称
- 注释(comments)
- 结构体
  - 如果结构体需要实现接口，可以通过声明可忽略接口变量，检查接口实现情况
  - 避免在公共结构体中切入类型
- 接口(interface)
  - 通过声明可忽略接口变量，检查接口实现情况
- 互斥器(mutex)
  - 默认值的`mutex`是有效的
  -
- 切片
  - 追加时优先确定切片容量
  - 指定切片容量
- 散列表
- 字符串
  - 避免字符串到字节的转换
- 通道
  - `channel` 的大小要么是1，要么是0
  -
- 关键字
  - 使用`defer`释放资源
  - 不要使用`panic`
- 错误
- 变量
  - 避免可变的全局变量
  -
- 常量
  - 常量的值尽量从1开始
- 时间
  - 使用`time.Time`表示瞬时时间
  - 使用`time.Duration`表示时间段
- 函数
  - 避免使用`init()`函数
  - 主函数退出方式

编程风格

1. 避免过长的行
2. 相似声明分组
3. 包引入分组
4. 文件按`struct` `const` `var`排序，函数按调用顺序排序
5. 通过优先处理特殊情况并尽早返回来减少嵌套
6. 减少不必要的`else`
7. 结构体中的嵌入字段和常规字段应该分开
8. 使用字段名称初始化结构体

## 命名

Go语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则：一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。大写字母和小写字母是不同的。

Go语言中类似if和switch的关键字有25个；关键字不能用于自定义名字，只能在特定语法结构中使用。

```text
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var
```

此外，还有大约30多个预定义的名字，比如int和true等，主要对应内建的常量、类型和函数。

```text
内建常量: true false iota nil

内建类型: int int8 int16 int32 int64
          uint uint8 uint16 uint32 uint64 uintptr
          float32 float64 complex128 complex64
          bool byte rune string error

内建函数: make len cap new append copy close delete
          complex real imag
          panic recover
```

在习惯上，Go语言程序员推荐使用 驼峰式 命名，当名字有几个单词组成的时优先使用大小写分隔，而不是优先用下划线分隔。

## 声明

Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。

## 注释

## 接口

### 接口的底层实现

接口的底层实现是由两个指针表示的，一个是指向**特定类型**的指针，另一个是指向**数据**的指针，对于数据指针，如果存储的数据是指针则直接存储，否则存储该数据值得指针。

通常不需要指向接口类型的指针，接口应该作为值进行传递，这样的传递过程中，传递的底层数据实际是指针。

如果希望接口方法修改基础数据，则必须使用指针传递，即将对象指针赋值给接口变量。

可以通过声明接口变量，编译期间验证结构体的对该接口实现情况的检查。

```go
package style

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

type Modifier interface {
    change()
}

type S1 struct {
    Value int
}

func (s S1) change() {
    s.Value = s.Value + 1
}

type S2 struct {
    Value int
}

func (s *S2) change() {
    s.Value = s.Value + 1
}

func TestPointerToInterface(t *testing.T) {
    var s1 = S1{Value: 1}
    var s2 = &S2{Value: 1}

    var m1 Modifier = s1
    var m2 Modifier = s2

    m1.change()
    m2.change()

    var b = s1.Value == s2.Value
    fmt.Printf(
        "S1 Value = %d, S2 Value = %d, (S1 == S2) is %v",
        s1.Value,
        s2.Value,
        b,
    )
    assert.Equal(t, s1.Value, s2.Value)
}
```

### 互斥器使用方式

`sync.Mutex`和`sync.ReMutex`默认值是有效的，所以指向其指针是不必要的，所以无须使用`new`关键字创建指针。

```go
package main

import "sync"

func bad() {
    mu := new(sync.Mutex)
    mu.Lock()
    defer mu.Unlock()
    //...
}

func good() {
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock()
    //...
}
```

当结构体需要支持线程安全时，`sync.Mutex`和`sync.ReMutex`应该作为结构体的非指针字段，并且不应该将其直接嵌入到结构体中，因为`mutex`及其方法是该结构体的实现细节，应该对其调用者不可见。

```go
package main

import "sync"

type SMap struct {
    //mu *sync.Mutex
    // sync.Mutex
    mu sync.Mutex

    data map[string]string
}

func NewSMap() *SMap {
    return &SMap{
        data: make(map[string]string),
    }
}

func (m *SMap) Get(k string) string {
    m.mu.Lock()
    defer m.mu.Unlock()

    return m.data[k]
}
```

### 资源释放

使用`defer`释放资源，如文件、锁、数据库连接等，使用其开销很小，并且可以增加代码的可读性。

```go
package main

import "sync"

type Counter struct {
    mu      sync.Mutex
    counter int
}

func badAdd() int {
    c := Counter{}
    c.mu.Lock()

    if c.counter > 10 {
        c.mu.Unlock()
        return c.counter
    }

    c.counter++
    ret := c.counter
    c.mu.Unlock()
    return ret
}

func goodAdd() int {
    c := Counter{}
    c.mu.Lock()
    defer c.mu.Unlock()

    if c.counter > 10 {
        return c.counter
    }
    c.counter++
    return c.counter
}
```

### 常量的使用

### 时间操作

---

# 设计模式

## 创建型

## 结构型

## 行为型

---

# 经典算法

---

# 面试题

1. 描述`Go`语言中指针和`C/C++`中有什么区别，分析该设计的优缺点。
2. 描述`Go`语言中`Slice`的实现原理以及元素访问、追加、扩容过程。
3. 描述`Go`语言中`Map`的实现原理以及元素访问、写入、删除、扩容过程。
4. 描述`Go`语言中关键字`make`和`new`各自的作用，对比其异同。
5. 描述`Go`语言中关键字`select`的作用，列举使用`select`时可能出现的现象并分析其数据结构和实现原理。
6. 描述`Go`语言中关键字`defer`的作用，列举使用`defer`时可能出现的现象并分析其数据结构和实现原理。
7. 描述`Go`语言中`context.Context`的作用，并分析其设计原理。
8. 描述`Go`语言中`sync.WaitGroup`的基本使用和设计实现。