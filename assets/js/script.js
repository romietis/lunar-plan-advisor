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
function calculateBestPlans() {
    var amountText = document.getElementById("amount").value.trim();
    var amount = parseFloat(amountText);
    if (isNaN(amount)) {
        alert("Please enter a valid amount.");
        return;
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

    displayBestPlans(bestPlans);
}

// Function to display the best plans
function displayBestPlans(plans) {
    var container = document.getElementById("best-plans");
    container.innerHTML = "";

    if (plans.length > 0) {
        var header = document.createElement("h3");
        header.textContent = "Best plans for you:";
        container.appendChild(header);

        plans.forEach(function (plan) {
            var div = document.createElement("div");
            div.className = "plan";

            var name = document.createElement("p");
            name.textContent = "Plan: " + plan.name;
            div.appendChild(name);

            var annualIncome = document.createElement("p");
            annualIncome.textContent = "Annual Interest: " + plan.annualInterest.toFixed(2);
            div.appendChild(annualIncome);

            var monthlyIncome = document.createElement("p");
            monthlyIncome.textContent = "Monthly Interest: " + plan.monthlyInterest.toFixed(2);
            div.appendChild(monthlyIncome);

            var annualProfit = document.createElement("p");
            annualProfit.textContent = "Annual Interest after fees: " + plan.annualInterestProfit.toFixed(2);
            div.appendChild(annualProfit);

            var monthlyProfit = document.createElement("p");
            monthlyProfit.textContent = "Monthly Interest after fee: " + plan.monthlyInterestProfit.toFixed(2);
            div.appendChild(monthlyProfit);

            container.appendChild(div);
        });
    }
}
