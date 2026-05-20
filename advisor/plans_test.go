package advisor

import (
	"fmt"
	"math"
	"testing"
)

func TestCalculatePlans_NegativeBalance(t *testing.T) {
	plans := PlansConfig{}
	_, err := plans.CalculatePlans(-1)
	if err == nil || err.Error() != "balance can't be negative" {
		t.Errorf("Expected error for negative balance, got %v", err)
	}
}

func TestCalculatePlansNoPlans(t *testing.T) {
	plans := PlansConfig{}
	bestPlans, _ := plans.CalculatePlans(1000)
	if len(bestPlans) != 0 {
		t.Errorf("Expected no plans, got %v", bestPlans)
	}
}

func TestCalculatePlansSinglePlan(t *testing.T) {
	plan := PlanConfig{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plans := PlansConfig{Plans: []PlanConfig{plan}}
	bestPlans, _ := plans.CalculatePlans(1000)
	if len(bestPlans) != 1 || bestPlans[0].Name != "Plan 1" {
		t.Errorf("Expected Plan 1, got %v", bestPlans)
	}
}

func TestCalculatePlansMultiplePlans(t *testing.T) {
	plan1 := PlanConfig{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plan2 := PlanConfig{Name: "Plan 2", AnnualInterestRate: 6, Fee: 1}
	plans := PlansConfig{Plans: []PlanConfig{plan1, plan2}}
	bestPlans, _ := plans.CalculatePlans(1000)
	if len(bestPlans) != 1 || bestPlans[0].Name != "Plan 2" {
		t.Errorf("Expected Plan 2, got %v", bestPlans)
	}
}

func TestCalculatePlansMultipleBestPlans(t *testing.T) {
	plan1 := PlanConfig{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plan2 := PlanConfig{Name: "Plan 2", AnnualInterestRate: 5, Fee: 1}
	plans := PlansConfig{Plans: []PlanConfig{plan1, plan2}}
	bestPlans, _ := plans.CalculatePlans(1000)
	if len(bestPlans) != 2 || bestPlans[0].Name != "Plan 1" || bestPlans[1].Name != "Plan 2" {
		t.Errorf("Expected Plan 1 and Plan 2, got %v", bestPlans)
	}
}

func TestCalculatePlansEffectiveBalanceIsCap(t *testing.T) {
	plan1 := PlanConfig{Name: "Plan 1", AnnualInterestRate: 1.5, Fee: 0, Cap: 1000}
	plan2 := PlanConfig{Name: "Plan 2", AnnualInterestRate: 2.25, Fee: 119, Cap: 0}
	plans := PlansConfig{Plans: []PlanConfig{plan1, plan2}}
	bestPlans, _ := plans.CalculatePlans(1001)
	if len(bestPlans) != 1 || bestPlans[0].Name != "Plan 1" {
		t.Errorf("Expected Plan 1, got %v", plans)
	}
}

func TestCalculatePlansCompound(t *testing.T) {
	planConfig := []PlanConfig{
		{Name: "Light", AnnualInterestRate: 1.25, Fee: 0.0, Cap: 100000},
		{Name: "Standard", AnnualInterestRate: 1.5, Fee: 29.0, Cap: 100000},
		{Name: "Plus", AnnualInterestRate: 1.75, Fee: 69.0, Cap: 0},
		{Name: "Unlimited", AnnualInterestRate: 2.25, Fee: 139.0, Cap: 0},
	}
	plans := PlansConfig{Plans: planConfig}

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

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		plans   PlansConfig
		wantErr bool
	}{
		{"empty list", PlansConfig{}, true},
		{"empty name", PlansConfig{Plans: []PlanConfig{{Name: "", AnnualInterestRate: 1}}}, true},
		{"negative rate", PlansConfig{Plans: []PlanConfig{{Name: "A", AnnualInterestRate: -1}}}, true},
		{"negative fee", PlansConfig{Plans: []PlanConfig{{Name: "A", Fee: -1}}}, true},
		{"negative cap", PlansConfig{Plans: []PlanConfig{{Name: "A", Cap: -1}}}, true},
		{"NaN rate", PlansConfig{Plans: []PlanConfig{{Name: "A", AnnualInterestRate: math.NaN()}}}, true},
		{"Inf fee", PlansConfig{Plans: []PlanConfig{{Name: "A", Fee: math.Inf(1)}}}, true},
		{"valid", PlansConfig{Plans: []PlanConfig{{Name: "A", AnnualInterestRate: 1, Fee: 0, Cap: 0}}}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.plans.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() err = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func ExamplePlansConfig_CalculatePlans() {
	plan1 := PlanConfig{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plan2 := PlanConfig{Name: "Plan 2", AnnualInterestRate: 6, Fee: 1}
	plans := PlansConfig{Plans: []PlanConfig{plan1, plan2}}
	bestPlans, _ := plans.CalculatePlans(1000)
	fmt.Println(bestPlans[0].Name)
	// Output: Plan 2
}

func ExamplePlansConfig_CalculatePlans_two_best_plans() {
	plan1 := PlanConfig{Name: "Plan 1", AnnualInterestRate: 5, Fee: 1}
	plan2 := PlanConfig{Name: "Plan 2", AnnualInterestRate: 5, Fee: 1}
	plans := PlansConfig{Plans: []PlanConfig{plan1, plan2}}
	bestPlans, _ := plans.CalculatePlans(1000)
	fmt.Println(len(bestPlans))
	// Output: 2
}
