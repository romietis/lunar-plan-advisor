package advisor

import (
	"errors"
	"math"
)

// PlanConfig is the user-configurable part of a plan: the inputs that define it.
type PlanConfig struct {
	Name               string  `json:"name"`
	AnnualInterestRate float64 `json:"annualInterestRate"`
	Fee                float64 `json:"fee"`
	Cap                float64 `json:"cap"`
}

// Plan is a configured plan together with the values computed for a given balance.
// PlanConfig is embedded so the JSON output keeps a flat shape.
type Plan struct {
	PlanConfig
	AnnualFee              float64 `json:"annualFee"`
	AnnualInterest         float64 `json:"annualInterest"`
	MonthlyInterest        float64 `json:"monthlyInterest"`
	AnnualInterestProfit   float64 `json:"annualInterestProfit"`
	MonthlyInterestProfit  float64 `json:"monthlyInterestProfit"`
	AnnualCompoundInterest float64 `json:"annualCompoundInterest"`
	AnnualCompoundProfit   float64 `json:"annualCompoundProfit"`
}

// PlansConfig is a collection of plan configurations — the input to CalculatePlans.
type PlansConfig struct {
	Plans []PlanConfig `json:"plans"`
}

// Plans is the calculated result returned by CalculatePlans.
type Plans struct {
	Plans []Plan `json:"plans"`
}

// CalculatePlans calculates the best plan(s) for the given balance against the
// plan configurations. Multiple plans can share the same maximum profit.
func (pc *PlansConfig) CalculatePlans(balance float64) ([]Plan, error) {
	if balance < 0 {
		return nil, errors.New("balance can't be negative")
	}
	var bestPlans []Plan
	maxProfit := math.SmallestNonzeroFloat64

	for _, cfg := range pc.Plans {
		plan := Plan{PlanConfig: cfg}

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

// Validate checks that the plan configuration is well-formed: non-empty list,
// non-empty names, and finite, non-negative numeric fields.
func (pc *PlansConfig) Validate() error {
	if len(pc.Plans) == 0 {
		return errors.New("plans list can't be empty")
	}
	for _, p := range pc.Plans {
		if p.Name == "" {
			return errors.New("plan name can't be empty")
		}
		for _, f := range []float64{p.AnnualInterestRate, p.Fee, p.Cap} {
			if math.IsNaN(f) || math.IsInf(f, 0) {
				return errors.New("plan numeric fields must be finite")
			}
		}
		if p.AnnualInterestRate < 0 || p.Fee < 0 || p.Cap < 0 {
			return errors.New("plan numeric fields can't be negative")
		}
	}
	return nil
}
