package strategy

import "fmt"

// 定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则。

type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

type PaymentContext struct {
	Name   string
	CardId string
	Money  int
}

type Payment struct {
	ctx      *PaymentContext
	strategy PaymentStrategy
}

func newPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		ctx: &PaymentContext{
			Name:   name,
			CardId: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p Payment) Pay() {
	p.strategy.Pay(p.ctx)
}

type Cash struct {
}

func (c Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

var _ PaymentStrategy = &Cash{}

type Bank struct {
}

func (b *Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account", ctx.Money, ctx.Name)
}

var _ PaymentStrategy = &Bank{}
