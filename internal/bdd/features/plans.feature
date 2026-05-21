Feature: advice plan
    In order to make good financial decision
    As a Lunar customer
    I need to be able to chose most profitable plan

Scenario: input balance
    Given a blance of 1000.00 DKK
    When I send "POST" request to "/plans/best"
    Then the response code should be 200

Scenario: negative balance is rejected
    Given a blance of -100.00 DKK
    When I send "POST" request to "/plans/best"
    Then the response code should be 400
