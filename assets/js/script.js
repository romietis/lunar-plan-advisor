// Function to call localhost endpoint
function calculateBestPlans() {
    var balanceText = document.getElementById("balance").value.trim();
    var balance = parseFloat(balanceText);
    if (isNaN(balance)) {
        alert("Please enter a valid balance");
        return;
    }
    fetch("/plans?balance=" + balance)
    .then(response => response.json())
    .then(data => {
        displayBestPlans(data.plans);
    })
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
