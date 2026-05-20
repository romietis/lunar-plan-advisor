// Load the user's plan configuration, seeding from server defaults on first visit.
async function getPlans() {
    const cached = localStorage.getItem("plans");
    if (cached) return JSON.parse(cached);
    const fresh = await fetch("/plans").then(r => r.json());
    localStorage.setItem("plans", JSON.stringify(fresh));
    return fresh;
}

function importConfig(file) {
    const reader = new FileReader();
    reader.onload = e => {
        try {
            const parsed = JSON.parse(e.target.result);
            if (!Array.isArray(parsed.plans)) throw new Error("missing 'plans' array");
            localStorage.setItem("plans", JSON.stringify(parsed));
            renderCurrentConfig();
            alert("Config imported");
        } catch (err) {
            alert("Invalid config: " + err.message);
        }
    };
    reader.readAsText(file);
}

function exportConfig() {
    const data = localStorage.getItem("plans") || "{}";
    const a = document.createElement("a");
    a.href = URL.createObjectURL(new Blob([data], { type: "application/json" }));
    a.download = "plans.json";
    a.click();
    URL.revokeObjectURL(a.href);
}

async function resetConfig() {
    localStorage.removeItem("plans");
    await getPlans();
    renderCurrentConfig();
    alert("Config reset to defaults");
}

async function renderCurrentConfig() {
    const target = document.getElementById("current-config");
    if (!target) return;
    const cfg = await getPlans();
    target.value = JSON.stringify(cfg, null, 2);
}

function saveConfig() {
    const target = document.getElementById("current-config");
    try {
        const parsed = JSON.parse(target.value);
        if (!Array.isArray(parsed.plans)) throw new Error("missing 'plans' array");
        localStorage.setItem("plans", JSON.stringify(parsed));
        alert("Config saved");
    } catch (err) {
        alert("Invalid config: " + err.message);
    }
}

async function calculateBestPlans() {
    const balance = parseFloat(document.getElementById("balance").value.trim());
    if (isNaN(balance)) {
        alert("Please enter a valid balance");
        return;
    }
    const cfg = await getPlans();
    const res = await fetch("/plans/best", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ balance, plans: cfg.plans }),
    });
    if (!res.ok) {
        const err = await res.json().catch(() => ({}));
        alert("Calculation failed: " + (err.error || res.statusText));
        return;
    }
    const data = await res.json();
    displayBestPlans(data.plans);
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

window.addEventListener("DOMContentLoaded", renderCurrentConfig);
