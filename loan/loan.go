package loan

import (
	"math"
	"time"
)

type APR struct {
	value                        float64
	diff                         float64
	percentageOfAllowedThreshold float64
	overThreshold                bool
}

func (a *APR) Value() float64 {
	return a.value
}

func (a *APR) Diff() float64 {
	return a.diff
}

func (a *APR) PercentageOfAllowedThreshold() float64 {
	return a.percentageOfAllowedThreshold
}

func (a *APR) OverThreshold() bool {
	return a.overThreshold
}

func NewAPR(value float64, diff float64, percentageOfAllowedThreshold float64, overThreshold bool) *APR {
	return &APR{
		value:                        value,
		diff:                         diff,
		percentageOfAllowedThreshold: percentageOfAllowedThreshold,
		overThreshold:                overThreshold,
	}
}

type Loan struct {
	principal             float64
	payments              []Payment
	startDate             time.Time
	monthlyPayment        float64
	term                  int
	daysUntilFirstPayment int
	finalPayment          float64
	nominalRate           float64
	totalInterest         float64
	bisectionSimpleAPR    APR
	bisectionActualAPR    APR
	simpleAPR             float64
}

func (l *Loan) Principal() float64 {
	return l.principal
}

func (l *Loan) Payments() []Payment {
	return l.payments
}

func (l *Loan) StartDate() time.Time {
	return l.startDate
}

func (l *Loan) MonthlyPayment() float64 {
	return l.monthlyPayment
}

func (l *Loan) Term() int {
	return l.term
}

func (l *Loan) DaysUntilFirstPayment() int {
	return l.daysUntilFirstPayment
}

func (l *Loan) FinalPayment() float64 {
	return l.finalPayment
}

func (l *Loan) NominalRate() float64 {
	return l.nominalRate
}

func (l *Loan) TotalInterest() float64 {
	return l.totalInterest
}

func (l *Loan) BisectionSimpleAPR() *APR {
	return &l.bisectionSimpleAPR
}

func (l *Loan) BisectionActualAPR() *APR {
	return &l.bisectionActualAPR
}

func (l *Loan) SimpleAPR() float64 {
	return l.simpleAPR
}

func monthlyPayment(principal float64, rate float64, term int) float64 {
	payment := principal * rate / (1 - math.Pow(1+rate, -float64(term)))
	return payment
}

type NextPaymentDate func(currentDate time.Time, startDayOfMonth int) time.Time

func calculatePayments(principal float64, startDate time.Time, daysUntilFirstPayment int, nominalRate float64, term int, finalPayment float64, nextPaymentDate NextPaymentDate) ([]Payment, float64) {
	payments := []Payment{}
	totalInterest := 0.0
	balance := principal
	currentPaymentDate := startDate.AddDate(0, 0, daysUntilFirstPayment)
	startDayOfMonth := currentPaymentDate.Day()
	currentPayment := monthlyPayment(balance, nominalRate/12, term)
	for i := 1; i <= term; i++ {
		interestPortion := balance * (nominalRate / 12)
		totalInterest += interestPortion

		principalPortion := currentPayment - interestPortion
		balance -= principalPortion

		// set current payment to final payment if last payment, otherwise monthly payment
		if i == term {
			payments = append(payments, *NewPayment(finalPayment, principalPortion, finalPayment-principalPortion, currentPaymentDate, balance))
		} else {
			payments = append(payments, *NewPayment(currentPayment, principalPortion, interestPortion, currentPaymentDate, balance))
		}

		currentPaymentDate = nextPaymentDate(currentPaymentDate, startDayOfMonth)
	}
	return payments, totalInterest
}

func nextPaymentDateWithEndOfMonth(currentDate time.Time, startDayOfMonth int) time.Time {
	year, month, _ := currentDate.Date()

	// get the number of days in the next month
	daysInNextMonth := time.Date(year, month+2, 0, 0, 0, 0, 0, time.UTC).Day()

	// if the day of the first payment is great than the number of days in the next month
	// set the day of the month to the last day of the month
	day := startDayOfMonth
	if day > daysInNextMonth {
		day = daysInNextMonth
	}

	return time.Date(year, month+1, day, 0, 0, 0, 0, time.UTC)
}

func simplePresentValue(principal float64, rate float64, term int, monthlyPmt float64, finalPayment float64, daysUntilFirstPayment int) F {
	return func(r float64) float64 {
		pv := -principal
		for i := 0; i < term-1; i++ { // notice the "-1", as the last month has a different payment
			pv += monthlyPmt / math.Pow(1+r/12, float64(i)+float64(daysUntilFirstPayment)/30.0)
		}
		// Add the irregular final payment
		pv += finalPayment / math.Pow(1+r/12, float64(term-1)+float64(daysUntilFirstPayment)/30.0)
		return pv
	}
}

func actualPresentValue(principal float64, rate float64, term int, payments []Payment, startDate time.Time) F {
	return func(r float64) float64 {
		pv := -principal

		for _, payment := range payments {
			startYear, startMonth, startDay := startDate.Date()
			paymentYear, paymentMonth, paymentDay := payment.Date().Date()

			daysDifference := (paymentYear-startYear)*360 +
				(int(paymentMonth)-int(startMonth))*30 +
				(paymentDay - startDay)

			pv += payment.Amount() / math.Pow(1+r/360, float64(daysDifference)) // Adjust denominator for 30/360 day count
		}

		return pv
	}
}

func simpleAPR(principal float64, monthlyPmt float64, nMonths int, rate float64) float64 {
	// use the simple interest formula for this calculation
	simpleInterest := principal * rate * float64(nMonths) / 12.0
	daysInLoanTerm := float64(nMonths) / 12 * 365.0
	return (simpleInterest / principal) * (1 / daysInLoanTerm) * 365.0
}

type F func(float64) float64

func bisect(lowerBound float64, upperBound float64, tolerance float64, f F) float64 {
	for upperBound-lowerBound > tolerance {
		midPoint := (lowerBound + upperBound) / 2.0

		if f(midPoint)*f(lowerBound) < 0 {
			upperBound = midPoint
		} else {
			lowerBound = midPoint
		}
	}

	return (lowerBound + upperBound) / 2.0
}

func NewLoan(loanAmount float64, nominalRate float64, term int, startDate time.Time, daysUntilFirstPayment int, finalPayment float64, lowerAPRBound float64, upperAPRBound float64, aprThreshold float64) *Loan {
	monthlyPayment := monthlyPayment(loanAmount, nominalRate/12, term)
	if finalPayment == 0 {
		finalPayment = monthlyPayment
	}
	payments, totalInterest := calculatePayments(loanAmount, startDate, daysUntilFirstPayment, nominalRate, term, finalPayment, nextPaymentDateWithEndOfMonth)
	for _, payment := range payments {
		totalInterest += payment.InterestAmount()
	}

	simpleAPR := simpleAPR(loanAmount, monthlyPayment, term, nominalRate)
	bisectSimpleAPR := bisect(lowerAPRBound, upperAPRBound, aprThreshold, simplePresentValue(loanAmount, nominalRate, term, monthlyPayment, finalPayment, daysUntilFirstPayment))
	bisectSimpleAPRDiff := math.Abs(bisectSimpleAPR - simpleAPR)
	bisectActualAPR := bisect(lowerAPRBound, upperAPRBound, aprThreshold, actualPresentValue(loanAmount, nominalRate, term, payments, startDate))
	bisectActualAPRDiff := math.Abs(bisectActualAPR - simpleAPR)

	regZThreshold := 0.01 / 8.0

	return &Loan{
		principal:             loanAmount,
		payments:              payments,
		startDate:             startDate,
		monthlyPayment:        monthlyPayment,
		term:                  term,
		daysUntilFirstPayment: daysUntilFirstPayment,
		finalPayment:          finalPayment,
		totalInterest:         totalInterest,
		bisectionSimpleAPR:    *NewAPR(bisectSimpleAPR, bisectSimpleAPRDiff, bisectSimpleAPRDiff/(1.0/800), bisectSimpleAPRDiff > regZThreshold),
		bisectionActualAPR:    *NewAPR(bisectActualAPR, bisectActualAPRDiff, bisectActualAPRDiff/(1.0/800), bisectActualAPRDiff > regZThreshold),
		simpleAPR:             simpleAPR,
	}
}
