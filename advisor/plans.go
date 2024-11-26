package advisor

import (
	"errors"
	"math"
)

type Plan struct {
	Name                   string  `json:"name"`
	AnnualInterestRate     float64 `json:"annualInterestRate"`
	Fee                    float64 `json:"fee"`
	Cap                    float64 `json:"cap"`
	AnnualFee              float64 `json:"annualFee"`
	AnnualInterest         float64 `json:"annualInterest"`
	MonthlyInterest        float64 `json:"monthlyInterest"`
	AnnualInterestProfit   float64 `json:"annualInterestProfit"`
	MonthlyInterestProfit  float64 `json:"monthlyInterestProfit"`
	AnnualCompoundInterest float64 `json:"annualCompoundInterest"`
	AnnualCompoundProfit   float64 `json:"annualCompoundProfit"`
}

type Plans struct {
	Plans []Plan `json:"plans"`
}

// CalculatePlans calculates the best investment plans based on the given balance and plan configurations.
// It returns a slice of plans because multiple plans can have the same maximum profit.
func (plans *Plans) CalculatePlans(balance float64) ([]Plan, error) {
	if balance < 0 {
		return nil, errors.New("balance can't be negative")
	}
	var bestPlans []Plan
	maxProfit := math.SmallestNonzeroFloat64

	for _, plan := range plans.Plans {
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
		plan.AnnualCompoundInterest = effectiveBalance*math.Pow(1+(plan.AnnualInterestRate/100)/12, 12) - effectiveBalance
		plan.AnnualCompoundProfit = plan.AnnualCompoundInterest - plan.AnnualFee

		if plan.AnnualCompoundProfit > maxProfit {
			maxProfit = plan.AnnualCompoundProfit
			bestPlans = []Plan{plan}
		} else if plan.AnnualCompoundProfit == maxProfit {
			bestPlans = append(bestPlans, plan)
		}
	}
	return bestPlans, nil
}
