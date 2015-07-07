Then(/^I should see a QRCode$/) do
  expect(page).to have_selector('#totp-qr-code', visible: true)
end
