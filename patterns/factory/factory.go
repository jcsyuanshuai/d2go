package factory

/**
go 设计模式
工厂
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
