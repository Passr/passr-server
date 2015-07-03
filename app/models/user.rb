class User < ActiveRecord::Base
  authenticates_with_sorcery!

  attr_accessor :password_confirmation

  validates :password, confirmation: true
  validates :password_confirmation, presence: {
    if: -> (obj) { obj.password.present? }
  }
end
