When(/^I fill in and submit the "(.*?)" form$/) do |form_name|
  form_method = {
    'sign up' => :sign_up_form,
    'sign in' => :sign_in_form
  }[form_name]

  raise "Unknown Form: #{form_name}" unless form_method

  send(form_method)
end

Given(/^I have a user account$/) do
  @user = FactoryGirl.create(:user)
end

Given(/^I am logged in$/) do
  visit(path_for('sign in'))

  sign_in_form
end

def sign_up_form
  fill_in 'Email', with: 'john.doe@passr.io'
  fill_in 'Password', with: 'passw0rd'
  fill_in 'Password confirmation', with: 'passw0rd'

  click_button 'Sign Up'
end

def sign_in_form
  fill_in 'Email', with: @user.email
  fill_in 'Password', with: 'password'

  click_button 'Sign In'
end
