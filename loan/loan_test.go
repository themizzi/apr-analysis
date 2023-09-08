package loan

import (
	"math"
	"testing"
	"time"
)

// doing this with all the floating points. never dealt with currency in go.

func TestMonthlyPayment(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)
	if math.Abs(loan.MonthlyPayment()-87.915887) > 0.000001 {
		t.Errorf("MonthlyPayment() = %f; want 87.915887", loan.MonthlyPayment())
	}
}

func TestTotalInterest(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)
	if math.Abs(loan.TotalInterest()-54.990647) > 0.000001 {
		t.Errorf("TotalInterest() = %f; want 54.990647", loan.TotalInterest())
	}
}

func TestPayments(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)
	payments := loan.Payments()
	if len(payments) != 12 {
		t.Errorf("len(Payments()) = %d; want 12", len(payments))
	}
	// test the properties of each payment
	for _, payment := range payments {
		if math.Abs(payment.Amount()-87.915887) > 0.000001 {
			t.Errorf("payment.Amount() = %f; want 87.915887", payment.Amount())
		}
		// TODO other props
	}
}

func TestSimpleAPR(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if math.Abs((loan.SimpleAPR())-0.10) > 0.000001 {
		t.Errorf("SimpleAPR() = %f; want 0.10", loan.SimpleAPR())
	}
}

func TestBisectionSimpleAPR(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if math.Abs((loan.BisectionSimpleAPR()).Value()-0.100006) > 0.000001 {
		t.Errorf("BisectionSimpleAPR() = %f; want 0.100006", loan.BisectionSimpleAPR().Value())
	}
}

func TestBisectionSimpleAPRPercentageOfAllowedThreshold(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if math.Abs(loan.BisectionSimpleAPR().PercentageOfAllowedThreshold()-0.004883) > 0.000001 {
		t.Errorf("BisectionSimpleAPR().PercentageOfAllowedThreshold() = %f; want 0.004883", loan.BisectionSimpleAPR().PercentageOfAllowedThreshold())
	}
}

func TestBisectionSimpleAPRDiff(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if math.Abs(loan.BisectionSimpleAPR().Diff()-0.000006) > 0.000001 {
		t.Errorf("BisectionSimpleAPR().Diff() = %f; want 0.000006", loan.BisectionSimpleAPR().Diff())
	}
}

func TestBisectionSimpleAPROverThreshold(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if loan.BisectionSimpleAPR().OverThreshold() {
		t.Errorf("BisectionSimpleAPR().OverThreshold() = %t; want false", loan.BisectionSimpleAPR().OverThreshold())
	}
}

func TestBisectinActualAPR(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if math.Abs((loan.BisectionActualAPR()).Value()-0.099884) > 0.000001 {
		t.Errorf("BisectionActualAPR() = %f; want 0.099884", loan.BisectionActualAPR().Value())
	}
}

func TestBisectionActualAPRPercetageOfAllowedThreshold(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if math.Abs(loan.BisectionActualAPR().PercentageOfAllowedThreshold()-0.092773) > 0.000001 {
		t.Errorf("BisectionActualAPR().PercentageOfAllowedThreshold() = %f; want 0.092773", loan.BisectionActualAPR().PercentageOfAllowedThreshold())
	}
}

func TestBisectionActualAPROverThreshold(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if loan.BisectionActualAPR().OverThreshold() {
		t.Errorf("BisectionActualAPR().OverThreshold() = %t; want false", loan.BisectionActualAPR().OverThreshold())
	}
}

func TestBisectionActualAPRDiff(t *testing.T) {
	loan := NewLoan(1000, 0.1, 12, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), 30, 0, 0.0, 0.5, 0.0001)

	if math.Abs(loan.BisectionActualAPR().Diff()-0.000116) > 0.000001 {
		t.Errorf("BisectionActualAPR().Diff() = %f; want 0.000116", loan.BisectionActualAPR().Diff())
	}
}
