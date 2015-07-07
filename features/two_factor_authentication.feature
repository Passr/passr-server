Feature: Two Factor Authentication
  As a user of Passr.io
  I want to enable two factor authentication
  So that my account is more secure

  Scenario: Enable Two Factor Authentication
    Given I have a user account
    And I am logged in
    And I am on the "dashboard" page
    When I click on "Enable Two Factor Authentication"
    Then I should see a QRCode
    When I scan the QRCode with my two factor authentication app
    And enter the code provided by the app
    And submit the form
    Then I should see "Two Factor Authentication Enabled"
