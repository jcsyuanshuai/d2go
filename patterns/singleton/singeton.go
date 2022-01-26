package singleton

import "sync"

//定义：一个类只允许创建一个对象(实例)，那这个类就是单例类，创建单例类的模式就是单例模式。
//用处：从业务概念上，有些数据在系统中只应该保存一份。

type Singleton struct {
}

var instance *Singleton = &Singleton{}

func GetInstance() *Singleton {
	return instance
}

var lazyInstance *Singleton

var once sync.Once

func GetLazyInstance() *Singleton {
	if lazyInstance == nil {
		once.Do(func() {
			lazyInstance = &Singleton{}
		})
	}
	return lazyInstance
}
