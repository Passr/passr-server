FactoryGirl.define do
  factory :user do
    email 'john.doe@passr.io'
    password 'password'
    password_confirmation 'password'
  end
end
