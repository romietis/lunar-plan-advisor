Feature: Get plans
  As a customer
  In order be well informed
  I need to be able to get advice on plans

  Scenario: GET request home
      # Given there are 4 plans
      When I send "GET" request to "/"
      Then the response code should be 200

  Scenario: POST request home
    When I send "POST" request to "/"
    Then the response code should be 400
