package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interestRate float32

	if balance < 0 {
		interestRate = 3.213
	} else if balance >= 0 && balance < 1000.0 {
		interestRate = 0.500
	} else if balance >= 1000.0 && balance < 5000 {
		interestRate = 1.621
	} else {
		interestRate = 2.475
	}

	return interestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	println(float64(InterestRate(balance)))
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	year := 0

	for i := balance; i < targetBalance; i += Interest(i) {
		year++
	}

	return year
}
