Feature: advice plan
    In order to make good financial decision
    As a Lunar customer
    I need to be able to chose most profitable plan

Scenario: input balance
    Given a blance of 1000.00 DKK
    When I send "GET" request to "/plans"
    Then the response code should be 200
    And the response should match json
