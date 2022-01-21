package strategy

func ExampleBank_Pay() {
	payment := newPayment("Ada", "", 123, &Cash{})
	payment.Pay()
}

func ExampleCash_Pay() {
	payment := newPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
}
