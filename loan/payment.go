package loan

import "time"

type Payment struct {
	amount          float64   // Total payment amount
	principalAmount float64   // Principal portion
	interestAmount  float64   // Interest portion
	date            time.Time // Date of the payment
	endingBalance   float64   // Remaining balance after payment
}

func (p *Payment) Amount() float64 {
	return p.amount
}

func (p *Payment) PrincipalAmount() float64 {
	return p.principalAmount
}

func (p *Payment) InterestAmount() float64 {
	return p.interestAmount
}

func (p *Payment) Date() time.Time {
	return p.date
}

func (p *Payment) EndingBalance() float64 {
	return p.endingBalance
}

func NewPayment(amount float64, principalAmount float64, interestAmount float64, date time.Time, endingBalance float64) *Payment {
	return &Payment{
		amount:          amount,
		principalAmount: principalAmount,
		interestAmount:  interestAmount,
		date:            date,
		endingBalance:   endingBalance,
	}
}
