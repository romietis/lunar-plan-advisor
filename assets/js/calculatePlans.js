// Define the Plan struct
function Plan(name, annualInterestRate, fee, cap = null) {
    this.name = name;
    this.annualInterestRate = annualInterestRate;
    this.fee = fee;
    this.cap = cap;

    // Initialize additional fields
    this.annualFee = fee * 12;
    this.annualInterest = 0;
    this.monthlyInterest = 0;
    this.annualInterestProfit =  0;
    this.monthlyInterestProfit = 0;
}

// Initialize instances of the Plan struct for each plan
var plans = [
    new Plan("Light", 1.5, 0.0, 100000.0),
    new Plan("Standard", 1.75, 29.0),
    new Plan("Plus", 2.0, 69.0),
    new Plan("Unlimited", 2.25, 119.0)
];

// Function to calculate the best plans based on the amount
function calculatePlans(amountText) {
    var amount = parseFloat(amountText);
    if (isNaN(amount)) {
        return null;
    }

    var bestPlans = [];
    var maxProfit = Number.NEGATIVE_INFINITY;

    plans.forEach(function (plan) {
        var effectiveAmount = plan.cap && plan.cap < amount ? plan.cap : amount;

        plan.annualInterest = (effectiveAmount * plan.annualInterestRate) / 100;
        plan.monthlyInterest = plan.annualInterest / 12;
        plan.annualInterestProfit = plan.annualInterest - plan.annualFee;
        plan.monthlyInterestProfit = plan.annualInterestProfit / 12;

        if (plan.annualInterestProfit > maxProfit) {
            maxProfit = plan.annualInterestProfit;
            bestPlans = [plan];
        } else if (plan.annualInterestProfit === maxProfit) {
            bestPlans.push(plan);
        }
    });

    return bestPlans;
}

module.exports = calculatePlans;
