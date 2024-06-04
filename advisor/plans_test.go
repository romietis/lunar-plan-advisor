package advisor

import (
	"testing"
)

func TestCalculatePlans_NegativeBalance(t *testing.T) {
	_, err := CalculatePlans(-1, nil)
	if err == nil || err.Error() != "balance can't be negative" {
		t.Errorf("Expected error for negative balance, got %v", err)
	}
}

func TestCalculatePlans_NoPlans(t *testing.T) {
	plans, _ := CalculatePlans(1000, nil)
	if len(plans) != 0 {
		t.Errorf("Expected no plans, got %v", plans)
	}
}

func TestCalculatePlans_SinglePlan(t *testing.T) {
	plan := Plan{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plans, _ := CalculatePlans(1000, []Plan{plan})
	if len(plans) != 1 || plans[0].Name != "Plan 1" {
		t.Errorf("Expected Plan 1, got %v", plans)
	}
}

func TestCalculatePlans_MultiplePlans(t *testing.T) {
	plan1 := Plan{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plan2 := Plan{Name: "Plan 2", AnnualInterestRate: 6, Fee: 1}
	plans, _ := CalculatePlans(1000, []Plan{plan1, plan2})
	if len(plans) != 1 || plans[0].Name != "Plan 2" {
		t.Errorf("Expected Plan 2, got %v", plans)
	}
}

func TestCalculatePlans_MultipleBestPlans(t *testing.T) {
	plan1 := Plan{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plan2 := Plan{Name: "Plan 2", AnnualInterestRate: 5, Fee: 1}
	plans, _ := CalculatePlans(1000, []Plan{plan1, plan2})
	if len(plans) != 2 || plans[0].Name != "Plan 1" || plans[1].Name != "Plan 2" {
		t.Errorf("Expected Plan 1 and Plan 2, got %v", plans)
	}
}

func TestCalculatePlans_EffectiveBalanceIsCap(t *testing.T) {
	plan1 := Plan{Name: "Plan 1", AnnualInterestRate: 1.5, Fee: 0, Cap: 1000}
	plan2 := Plan{Name: "Plan 2", AnnualInterestRate: 2.25, Fee: 119, Cap: 0}
	plans, _ := CalculatePlans(1001, []Plan{plan1, plan2})
	if len(plans) != 1 || plans[0].Name != "Plan 1" || plans[0].AnnualInterest != 15 {
		t.Errorf("Expected Plan 1 and annual interest 105, got %v", plans)
	}
}
