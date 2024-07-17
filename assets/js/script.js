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
            annualIncome.textContent = "Annual interest: " + plan.annualInterest.toFixed(2);
            div.appendChild(annualIncome);

            var monthlyIncome = document.createElement("p");
            monthlyIncome.textContent = "Monthly interest: " + plan.monthlyInterest.toFixed(2);
            div.appendChild(monthlyIncome);

            var annualProfit = document.createElement("p");
            annualProfit.textContent = "Annual interest profit after fees: " + plan.annualInterestProfit.toFixed(2);
            div.appendChild(annualProfit);

            var monthlyProfit = document.createElement("p");
            monthlyProfit.textContent = "Monthly interest after fee: " + plan.monthlyInterestProfit.toFixed(2);
            div.appendChild(monthlyProfit);
            
            var annualCompoundInterest = document.createElement("p");
            annualCompoundInterest.textContent = "Annual compound interest: " + plan.annualCompoundInterest.toFixed(2);
            div.appendChild(annualCompoundInterest);

            var annualCompoundProfit = document.createElement("p");
            annualCompoundProfit.textContent = "Annual compound interest profit after fees: " + plan.annualCompoundProfit.toFixed(2);
            div.appendChild(annualCompoundProfit);

            container.appendChild(div);
        });
    }
}
