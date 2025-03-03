package principles

import "fmt"

// LSP - Liskov Substitution Principle
// - If S is a subtype of T, then objects of type T may be replaced with objects of type S
// - Objects of a superclass shall be replaceable with objects of its subclasses without affecting the functionality of the program
// - It avoids the generalization of concepts that may not be needed in the future.
// - It makes the code maintainable and easier to upgrade.

// Only For Payment
type PaymentMethod interface {
	Pay(amount float32)
}

type RefundablePaymentMethod interface {
	PaymentMethod
	ProcessRefund(amount float32)
}

type CreditCard struct{} // Support Refund

func (cc CreditCard) Pay(amount float32) {
	fmt.Printf("Paying %.2f using Credit Card\n", amount)
}

func (cc CreditCard) ProcessRefund(amount float32) {
	fmt.Printf("Refunding %.2f using Credit Card\n", amount)
}

type Paypal struct{} // Support Refund

func (pp Paypal) Pay(amount float32) {
	fmt.Printf("Paying %.2f using Paypal\n", amount)
}

func (pp Paypal) ProcessRefund(amount float32) {
	fmt.Printf("Refunding %.2f using Paypal\n", amount)
}

type Cash struct{} // No Refund

func (c Cash) Pay(amount float32) {
	fmt.Printf("Paying %.2f using Cash\n", amount)
}

func IssueRefund(paymentMethod RefundablePaymentMethod, amount float32) {
	paymentMethod.ProcessRefund(amount)
}

func LSP() {
	fmt.Println("From LSP Package!!!")
	cc := CreditCard{}
	pp := Paypal{}
	cash := Cash{}
	cc.Pay(100)
	IssueRefund(cc, 50)
	pp.Pay(100)
	IssueRefund(pp, 50)
	cash.Pay(100)
	// IssueRefund(cash, 50) // This will not work as Cash does not support refund
}
