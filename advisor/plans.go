package advisor

import (
	"errors"
	"math"
)

type Plan struct {
	Name                  string  `json:"name"`
	AnnualInterestRate    float64 `json:"annualInterestRate"`
	Fee                   float64 `json:"fee"`
	Cap                   float64 `json:"cap"`
	AnnualFee             float64 `json:"annualFee"`
	AnnualInterest        float64 `json:"annualInterest"`
	MonthlyInterest       float64 `json:"monthlyInterest"`
	AnnualInterestProfit  float64 `json:"annualInterestProfit"`
	MonthlyInterestProfit float64 `json:"monthlyInterestProfit"`
}

func CalculatePlans(balance float64, planConfig []Plan) ([]Plan, error) {
	if balance < 0 {
		return nil, errors.New("balance can't be negative")
	}
	var bestPlans []Plan
	maxProfit := math.SmallestNonzeroFloat64

	for _, plan := range planConfig {
		var effectiveBalance float64
		if plan.Cap != 0 && plan.Cap < balance {
			effectiveBalance = plan.Cap
		} else {
			effectiveBalance = balance
		}

		plan.AnnualInterest = (effectiveBalance * plan.AnnualInterestRate) / 100
		plan.MonthlyInterest = plan.AnnualInterest / 12
		plan.AnnualFee = plan.Fee * 12
		plan.AnnualInterestProfit = plan.AnnualInterest - plan.AnnualFee
		plan.MonthlyInterestProfit = plan.AnnualInterestProfit / 12

		if plan.AnnualInterestProfit > maxProfit {
			maxProfit = plan.AnnualInterestProfit
			bestPlans = []Plan{plan}
		} else if plan.AnnualInterestProfit == maxProfit {
			bestPlans = append(bestPlans, plan)
		}
	}
	return bestPlans, nil
}
