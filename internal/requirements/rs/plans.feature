Feature: Get plan advice
  As a customer
  I want to know which plan yields highest annual interest profit based on my balance
  In order to make most of my money

  Scenario: GET request home
      # Given there are 4 plans
      When I send "GET" request to "/"
      Then the response code should be 200
