Feature: Get plans
    As a customer
    In order be well informed
    I need to be able to get advice on plans

    Scenario: Suggest plan
        Given there are 4 plans
        When I input balance
        Then I should get a suggestion