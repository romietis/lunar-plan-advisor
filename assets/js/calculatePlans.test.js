const assert = require('assert');
const calculatePlans = require('./calculatePlans');

// Test case for valid input
function testCalculatePlansWithValidAmount() {
    const amount = '1000'; // Assuming 10000 DKK
    const bestPlans = calculatePlans(amount);
    assert.ok(bestPlans); // Assert that the result is truthy
    assert.ok(Array.isArray(bestPlans)); // Assert that the result is an array

    assert.strictEqual(bestPlans.length, 1); // Assert that the result has a length of 1

    assert.strictEqual(bestPlans[0].name, 'Light');
    assert.strictEqual(bestPlans[0].annualInterest, 15);
    assert.strictEqual(bestPlans[0].monthlyInterest, 1.25);
    assert.strictEqual(bestPlans[0].annualInterestProfit, 15);
    assert.strictEqual(bestPlans[0].monthlyInterestProfit, 1.25);
}

// Test case for invalid input
function testCalculatePlansWithInvalidAmount() {
    const amount = 'invalid_amount';
    const bestPlans = calculatePlans(amount);
    assert.strictEqual(bestPlans, null); // Assert that the result is null
}

function testCalculatePlansWith105600Amount() {
    const amount = '105600';
    const bestPlans = calculatePlans(amount);
    assert.ok(bestPlans); // Assert that the result is truthy
    assert.ok(Array.isArray(bestPlans)); // Assert that the result is an array

    assert.strictEqual(bestPlans.length, 2); // Assert that the result has a length of 2
 
    assert.strictEqual(bestPlans[0].name, 'Light');
    assert.strictEqual(bestPlans[0].annualInterest, 1500);
    assert.strictEqual(bestPlans[0].monthlyInterest, 125);
    assert.strictEqual(bestPlans[0].annualInterestProfit, 1500);
    assert.strictEqual(bestPlans[0].monthlyInterestProfit, 125);

    assert.strictEqual(bestPlans[1].name, 'Standard');
    assert.strictEqual(bestPlans[1].annualInterest, 1848);
    assert.strictEqual(bestPlans[1].monthlyInterest, 154);
    assert.strictEqual(bestPlans[1].annualInterestProfit, 1500);
    assert.strictEqual(bestPlans[1].monthlyInterestProfit, 125);
}

function testCalculatePlansWith192000Amount() {
    const amount = '192000';
    const bestPlans = calculatePlans(amount);
    assert.ok(bestPlans); // Assert that the result is truthy
    assert.ok(Array.isArray(bestPlans)); // Assert that the result is an array

    assert.strictEqual(bestPlans.length, 2);
 
    assert.strictEqual(bestPlans[0].name, 'Standard');
    assert.strictEqual(bestPlans[0].annualInterest, 3360);
    assert.strictEqual(bestPlans[0].monthlyInterest, 280);
    assert.strictEqual(bestPlans[0].annualInterestProfit, 3012);
    assert.strictEqual(bestPlans[0].monthlyInterestProfit, 251);

    assert.strictEqual(bestPlans[1].name, 'Plus');
    assert.strictEqual(bestPlans[1].annualInterest, 3840);
    assert.strictEqual(bestPlans[1].monthlyInterest, 320);
    assert.strictEqual(bestPlans[1].annualInterestProfit, 3012);
    assert.strictEqual(bestPlans[1].monthlyInterestProfit, 251);
}

function testCalculatePlansWith240000Amount() {
    const amount = '240000';
    const bestPlans = calculatePlans(amount);
    assert.ok(bestPlans); // Assert that the result is truthy
    assert.ok(Array.isArray(bestPlans)); // Assert that the result is an array

    assert.strictEqual(bestPlans.length, 2);
 
    assert.strictEqual(bestPlans[0].name, 'Plus');
    assert.strictEqual(bestPlans[0].annualInterest, 4800);
    assert.strictEqual(bestPlans[0].monthlyInterest, 400);
    assert.strictEqual(bestPlans[0].annualInterestProfit, 3972);
    assert.strictEqual(bestPlans[0].monthlyInterestProfit, 331);

    assert.strictEqual(bestPlans[1].name, 'Unlimited');
    assert.strictEqual(bestPlans[1].annualInterest, 5400);
    assert.strictEqual(bestPlans[1].monthlyInterest, 450);
    assert.strictEqual(bestPlans[1].annualInterestProfit, 3972);
    assert.strictEqual(bestPlans[1].monthlyInterestProfit, 331);
}

// Run the tests
function runTests() {
    try {
        testCalculatePlansWithValidAmount();
        testCalculatePlansWithInvalidAmount();
        testCalculatePlansWith105600Amount();
        testCalculatePlansWith192000Amount();
        testCalculatePlansWith240000Amount();
        console.log('All tests passed successfully!');
    } catch (error) {
        console.error('One or more tests failed:', error);
    }
}

runTests();
