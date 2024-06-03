// Function to calculate the best plans based on the amount
function calculateBestPlans() {
    var amountText = document.getElementById("amount").value.trim();
    var amount = parseFloat(amountText);
    if (isNaN(amount)) {
        alert("Please enter a valid amount.");
        return;
    }
    bestPlans = calculatePlans(amount)
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

            fetch("/hello")
            .then(response => response.json())
            .then(data => console.log(data));

            container.appendChild(div);
        });
    }
}
