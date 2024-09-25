package advisor

import (
	"fmt"
	"testing"
)

func TestCalculatePlans_NegativeBalance(t *testing.T) {
	plans := Plans{}
	_, err := plans.CalculatePlans(-1)
	if err == nil || err.Error() != "balance can't be negative" {
		t.Errorf("Expected error for negative balance, got %v", err)
	}
}

func TestCalculatePlansNoPlans(t *testing.T) {
	plans := Plans{}
	bestPlans, _ := plans.CalculatePlans(1000)
	if len(bestPlans) != 0 {
		t.Errorf("Expected no plans, got %v", bestPlans)
	}
}

func TestCalculatePlansSinglePlan(t *testing.T) {
	plan := Plan{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plans := Plans{Plans: []Plan{plan}}
	bestPlans, _ := plans.CalculatePlans(1000)
	if len(bestPlans) != 1 || bestPlans[0].Name != "Plan 1" {
		t.Errorf("Expected Plan 1, got %v", bestPlans)
	}
}

func TestCalculatePlansMultiplePlans(t *testing.T) {
	plan1 := Plan{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plan2 := Plan{Name: "Plan 2", AnnualInterestRate: 6, Fee: 1}
	plans := Plans{Plans: []Plan{plan1, plan2}}
	bestPlans, _ := plans.CalculatePlans(1000)
	if len(bestPlans) != 1 || bestPlans[0].Name != "Plan 2" {
		t.Errorf("Expected Plan 2, got %v", bestPlans)
	}
}

func TestCalculatePlansMultipleBestPlans(t *testing.T) {
	plan1 := Plan{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plan2 := Plan{Name: "Plan 2", AnnualInterestRate: 5, Fee: 1}
	plans := Plans{Plans: []Plan{plan1, plan2}}
	bestPlans, _ := plans.CalculatePlans(1000)
	if len(bestPlans) != 2 || bestPlans[0].Name != "Plan 1" || bestPlans[1].Name != "Plan 2" {
		t.Errorf("Expected Plan 1 and Plan 2, got %v", bestPlans)
	}
}

func TestCalculatePlansEffectiveBalanceIsCap(t *testing.T) {
	plan1 := Plan{Name: "Plan 1", AnnualInterestRate: 1.5, Fee: 0, Cap: 1000}
	plan2 := Plan{Name: "Plan 2", AnnualInterestRate: 2.25, Fee: 119, Cap: 0}
	plans := Plans{Plans: []Plan{plan1, plan2}}
	bestPlans, _ := plans.CalculatePlans(1001)
	if len(bestPlans) != 1 || bestPlans[0].Name != "Plan 1" {
		t.Errorf("Expected Plan 1, got %v", plans)
	}
}

func TestCalculatePlansCompound(t *testing.T) {
	planConfig := []Plan{
		{Name: "Light", AnnualInterestRate: 1.25, Fee: 0.0, Cap: 100000},
		{Name: "Standard", AnnualInterestRate: 1.5, Fee: 29.0, Cap: 100000},
		{Name: "Plus", AnnualInterestRate: 1.75, Fee: 69.0, Cap: 0},
		{Name: "Unlimited", AnnualInterestRate: 2.25, Fee: 139.0, Cap: 0},
	}
	plans := Plans{Plans: planConfig}

	testCases := []struct {
		balance                  float64
		wantAnnualCompoundProfit float64
		wantName                 string
	}{
		{10000, 125.71863828854475, "Light"},
		{100000, 1257.1863828854548, "Light"},
		{160000, 1994.5678648465837, "Plus"},
		{200000, 2878.697516343789, "Unlimited"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.balance), func(t *testing.T) {
			got, err := plans.CalculatePlans(testCase.balance)
			if err != nil {
				t.Fatal(err)
			}
			if got[0].AnnualCompoundProfit != testCase.wantAnnualCompoundProfit && got[0].Name != testCase.wantName {
				t.Errorf("got %v, want %v", got, testCase)
			}
		})
	}

}
