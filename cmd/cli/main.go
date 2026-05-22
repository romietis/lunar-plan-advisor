package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/romietis/lunar-plan-advisor/v3/advisor"
)

func main() {
	// Check if balance argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: lunar-plan-advisor <balance>")
		fmt.Println("Example: lunar-plan-advisor 50000")
		os.Exit(1)
	}

	// Parse balance from command line argument
	balanceStr := os.Args[1]
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		fmt.Printf("Error: Invalid balance '%s'. Please provide a valid number.\n", balanceStr)
		os.Exit(1)
	}

	data, err := os.ReadFile("plans.json")
	if err != nil {
		fmt.Printf("Error reading plans.json: %v\n", err)
		os.Exit(1)
	}
	var planConfig advisor.PlansConfig
	if err := json.Unmarshal(data, &planConfig); err != nil {
		fmt.Printf("Error parsing plans.json: %v\n", err)
		os.Exit(1)
	}

	bestPlans, err := planConfig.CalculatePlans(balance)
	if err != nil {
		fmt.Printf("Error calculating plans: %v\n", err)
		os.Exit(1)
	}

	// Output results
	fmt.Printf("Best investment plan(s) for balance: %.2f\n", balance)
	fmt.Println("=" + fmt.Sprintf("%*s", 50, "="))

	for _, plan := range bestPlans {
		fmt.Printf("\nPlan: %s\n", plan.Name)
		fmt.Printf("Annual Interest Rate: %.2f%%\n", plan.AnnualInterestRate)
		fmt.Printf("Monthly Fee: %.2f\n", plan.Fee)
		fmt.Printf("Annual Fee: %.2f\n", plan.AnnualFee)
		fmt.Printf("Cap: %.2f\n", plan.Cap)
		fmt.Printf("Annual Interest: %.2f\n", plan.AnnualInterest)
		fmt.Printf("Monthly Interest: %.2f\n", plan.MonthlyInterest)
		fmt.Printf("Annual Interest Profit: %.2f\n", plan.AnnualInterestProfit)
		fmt.Printf("Monthly Interest Profit: %.2f\n", plan.MonthlyInterestProfit)
		fmt.Printf("Annual Compound Interest: %.2f\n", plan.AnnualCompoundInterest)
		fmt.Printf("Annual Compound Profit: %.2f\n", plan.AnnualCompoundProfit)
	}
}
