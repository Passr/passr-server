When(/^I visit the "(.*?)" page$/) do |page_name|
  visit(path_for(page_name))
end

Given(/^I am on the "(.*?)" page$/) do |page_name|
  visit(path_for(page_name))
end

When(/^I follow the "(.*?)" link$/) do |link_text|
  click_link(link_text)
end

Then(/^I should see "(.*?)"$/) do |text|
  expect(page).to have_text(text)
end

Then(/^I should not see "(.*?)"$/) do |text|
  expect(page).to have_no_text(text)
end

def path_for(page_name)
  {
    'sign up' => '/sign_up',
    'sign in' => '/sign_in',
    'home' => '/'
  }[page_name].tap do |path|
    raise "Unknown Page/Path #{page_name}" unless path
  end
end
