package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"apr/loan"
)

func main() {
	app := &cli.App{
		Name:  "Loan APR Calculator",
		Usage: "Calculate the APR and generate an amortization schedule for a loan",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:     "loan-amount",
				Usage:    "The total amount financed",
				Required: true,
			},
			&cli.Float64Flag{
				Name:     "nominal-rate",
				Usage:    "The nominal interest rate (e.g., 0.10 for 10%)",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "term",
				Usage:    "The term of the loan in months",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "start-date",
				Usage:    "The start date of the loan (YYYY-MM-DD format)",
				Required: true,
			},
			&cli.Float64Flag{
				Name:  "final-payment",
				Usage: "The final payment of the loan",
			},
			&cli.IntFlag{
				Name:  "days-until-first-payment",
				Usage: "The number of days until the first payment is due",
				Value: 30,
			},
		},
		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func action(c *cli.Context) error {
	loanAmount := c.Float64("loan-amount")
	// create new Loan
	startDate, err := time.Parse("2006-01-02", c.String("start-date"))
	if err != nil {
		return err
	}
	theLoan := loan.NewLoan(loanAmount, c.Float64("nominal-rate"), c.Int("term"), startDate, c.Int("days-until-first-payment"), c.Float64("final-payment"), 0, 0.5, 1e-6)

	fmt.Println("\033[1mAmortization Schedule\033[0m")
	fmt.Println("Month | Payment Date | Payment Amount   | Principal Amount | Interest Amount  | Balance")
	fmt.Println("------------------------------------------------------------------------------------------------")

	// print payments
	for i, payment := range theLoan.Payments() {
		fmt.Printf("%5d | %s   | $%15.2f | $%15.2f | $%15.2f | $%15.2f\n", i+1, payment.Date().Format("2006-01-02"), payment.Amount(), payment.PrincipalAmount(), payment.InterestAmount(), payment.EndingBalance())
	}

	fmt.Printf("\n\033[1mTotal Interest Paid:\033[0m $%.2f\n", theLoan.TotalInterest())
	fmt.Printf("\033[1mSimple APR:\033[0m %f%%\n", theLoan.SimpleAPR()*100.0)
	printAPR := func(apr *loan.APR, name string) {
		// print name in bold
		fmt.Printf("\n\033[1m%s\033[0m\n", name)
		fmt.Printf("APR: %51.8f%%\n", apr.Value()*100.0)
		fmt.Printf("Difference from simple APR: %28.8f%%\n", apr.Diff()*100.0)
		fmt.Printf("Percentage of tolerance level of 1/8 of 1%%: %12.8f%%\n", apr.PercentageOfAllowedThreshold()*100.0)
		var withinTolerance = "YES"
		if apr.OverThreshold() {
			withinTolerance = "NO"
		}
		fmt.Printf("Difference is within tolerance of 1/8 of 1%%:          %v\n", withinTolerance)
	}
	printAPR(theLoan.BisectionSimpleAPR(), "Bisection Method (Simple Formula)")
	printAPR(theLoan.BisectionActualAPR(), "BiSection Method (Actual Payments)")

	return nil
}
