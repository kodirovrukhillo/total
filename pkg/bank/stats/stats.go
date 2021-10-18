package stats

import (
	"github.com/kodirovrukhillo/total/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	var average types.Money
	for _, payment := range payments {
		if payment.Amount > 0 {
			average += payment.Amount

		}

	}
	average1 := average / types.Money(len(payments))
	return average1
}

func Total(payments []types.Payment, category types.Category) types.Money {
	var total types.Money

	for _, payment := range payments {
		if payment.Category == category && payment.Amount > 0 {
			total += payment.Amount
		}
	}
	return total
}
