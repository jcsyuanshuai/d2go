package factory

/**
go 设计模式 工厂模式

简单工厂：由于 Go 本身是没有构造函数的，一般而言我们采用 NewName
的方式创建对象/接口，当它返回的是接口的时候，其实就是简单工厂模式。

工厂方法：当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，
而是要组合其他类对象，做各种初始化操作的时候，推荐使用工厂方法模式，
将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂。
*/

type Operator interface {
	SetX(x int)
	SetY(y int)
	Result() int
}

type Factory interface {
	Create() Operator
}

type AbstractOperator struct {
	x int
	y int
}

func (a *AbstractOperator) SetX(x int) {
	a.x = x
}

func (a *AbstractOperator) SetY(y int) {
	a.y = y
}

//=================

type PlusOperator struct {
	*AbstractOperator
}

func (p PlusOperator) Result() int {
	return p.x + p.y
}

var _ Operator = &PlusOperator{}

type PlusFactory struct {
}

func (p PlusFactory) Create() Operator {
	return &PlusOperator{
		AbstractOperator: &AbstractOperator{},
	}
}

var _ Factory = &PlusFactory{}

//=================

type MinusOperator struct {
	AbstractOperator
}

func (m MinusOperator) Result() int {
	return m.x - m.y
}

var _ Operator = PlusOperator{}

type MinusFactory struct {
}

func (m MinusFactory) Create() Operator {
	return &MinusOperator{
		AbstractOperator: AbstractOperator{},
	}
}

var _ Factory = &MinusFactory{}
