# APR Analysis

## Usage

```bash
❯ go run . 

NAME:
   Loan APR Calculator - Calculate the APR and generate an amortization schedule for a loan

USAGE:
   Loan APR Calculator [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --loan-amount value               The total amount financed (default: 0)
   --nominal-rate value              The nominal interest rate (e.g., 0.10 for 10%) (default: 0)
   --term value                      The term of the loan in months (default: 0)
   --start-date value                The start date of the loan (YYYY-MM-DD format)
   --final-payment value             The final payment of the loan (default: 0)
   --days-until-first-payment value  The number of days until the first payment is due (default: 30)
   --help, -h                        show help
```

## Example

```bash
❯ go run . --loan-amount 10000 --nominal-rate 0.10 --term 36 --start-date 2022-01-01 --final-payment 322.67
Amortization Schedule
Month | Payment Date | Payment Amount   | Principal Amount | Interest Amount  | Balance
------------------------------------------------------------------------------------------------
    1 | 2022-01-31   | $         322.67 | $         239.34 | $          83.33 | $        9760.66
    2 | 2022-02-28   | $         322.67 | $         241.33 | $          81.34 | $        9519.33
    3 | 2022-03-31   | $         322.67 | $         243.34 | $          79.33 | $        9275.98
    4 | 2022-04-30   | $         322.67 | $         245.37 | $          77.30 | $        9030.61
    5 | 2022-05-31   | $         322.67 | $         247.42 | $          75.26 | $        8783.20
    6 | 2022-06-30   | $         322.67 | $         249.48 | $          73.19 | $        8533.72
    7 | 2022-07-31   | $         322.67 | $         251.56 | $          71.11 | $        8282.16
    8 | 2022-08-31   | $         322.67 | $         253.65 | $          69.02 | $        8028.51
    9 | 2022-09-30   | $         322.67 | $         255.77 | $          66.90 | $        7772.74
   10 | 2022-10-31   | $         322.67 | $         257.90 | $          64.77 | $        7514.84
   11 | 2022-11-30   | $         322.67 | $         260.05 | $          62.62 | $        7254.79
   12 | 2022-12-31   | $         322.67 | $         262.22 | $          60.46 | $        6992.58
   13 | 2023-01-31   | $         322.67 | $         264.40 | $          58.27 | $        6728.17
   14 | 2023-02-28   | $         322.67 | $         266.60 | $          56.07 | $        6461.57
   15 | 2023-03-31   | $         322.67 | $         268.83 | $          53.85 | $        6192.75
   16 | 2023-04-30   | $         322.67 | $         271.07 | $          51.61 | $        5921.68
   17 | 2023-05-31   | $         322.67 | $         273.32 | $          49.35 | $        5648.36
   18 | 2023-06-30   | $         322.67 | $         275.60 | $          47.07 | $        5372.75
   19 | 2023-07-31   | $         322.67 | $         277.90 | $          44.77 | $        5094.85
   20 | 2023-08-31   | $         322.67 | $         280.21 | $          42.46 | $        4814.64
   21 | 2023-09-30   | $         322.67 | $         282.55 | $          40.12 | $        4532.09
   22 | 2023-10-31   | $         322.67 | $         284.90 | $          37.77 | $        4247.19
   23 | 2023-11-30   | $         322.67 | $         287.28 | $          35.39 | $        3959.91
   24 | 2023-12-31   | $         322.67 | $         289.67 | $          33.00 | $        3670.23
   25 | 2024-01-31   | $         322.67 | $         292.09 | $          30.59 | $        3378.15
   26 | 2024-02-29   | $         322.67 | $         294.52 | $          28.15 | $        3083.63
   27 | 2024-03-31   | $         322.67 | $         296.97 | $          25.70 | $        2786.65
   28 | 2024-04-30   | $         322.67 | $         299.45 | $          23.22 | $        2487.20
   29 | 2024-05-31   | $         322.67 | $         301.95 | $          20.73 | $        2185.26
   30 | 2024-06-30   | $         322.67 | $         304.46 | $          18.21 | $        1880.80
   31 | 2024-07-31   | $         322.67 | $         307.00 | $          15.67 | $        1573.80
   32 | 2024-08-31   | $         322.67 | $         309.56 | $          13.11 | $        1264.24
   33 | 2024-09-30   | $         322.67 | $         312.14 | $          10.54 | $         952.10
   34 | 2024-10-31   | $         322.67 | $         314.74 | $           7.93 | $         637.37
   35 | 2024-11-30   | $         322.67 | $         317.36 | $           5.31 | $         320.01
   36 | 2024-12-31   | $         322.67 | $         320.01 | $           2.66 | $          -0.00

Total Interest Paid: $3232.37
Simple APR: 10.000000%

Bisection Method (Simple Formula)
APR:                                          9.99999046%
Difference from simple APR:                   0.00000954%
Percentage of tolerance level of 1/8 of 1%:   0.00762939%
Difference is within tolerance of 1/8 of 1%:          YES

BiSection Method (Actual Payments)
APR:                                          9.97061729%
Difference from simple APR:                   0.02938271%
Percentage of tolerance level of 1/8 of 1%:  23.50616455%
Difference is within tolerance of 1/8 of 1%:          YES
```