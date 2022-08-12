# keyword

## make vs. new

两者都是go中在堆空间分配内存的关键字，其中`new`接受一个参数，该参数是一个类型，分配好内存后，返回指向该类型的指针，同时将分配的内存置为该类型默认值。`make`关键字用于`chan`、`map`、`slice`的内存创建，接受2个或以上的参数，并且返回类型就是这三个类型本身。

## defer

1. 函数中 `defer` 和 `return` 那个先执行？
2. 多个`defer`的执行顺序是什么？
3. `defer`和`panic`那个先执行？
4. `defer`函数中包含子函数？

## select

## channel

### CSP

CSP，即不通过共享内存来通信，而通过通信来实现内存共享。


## interface

