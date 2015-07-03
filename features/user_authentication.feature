Feature: User Authentication
  In order to sign up and access private user pages
  As a visitor to the site
  I want to be able to sign up and authenticate

  Scenario: New User Creation
    Given I am on the "home" page
    When I follow the "Sign Up" link
    And I fill in and submit the "sign up" form
    Then I should see "Dashboard"
    And I should not see "Sign Up"

  Scenario: User Sign In
    Given I have a user account
    And I am on the "home" page
    When I follow the "Sign In" link
    And I fill in and submit the "sign in" form
    Then I should see "Dashboard"
