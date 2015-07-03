require 'rails_helper'

describe User, type: :model do
  describe 'password confirmation' do
    let(:user) { User.new(email: 'john.doe@example.com', password: 'passw0rd') }

    it 'is invalid without password confirmation' do
      expect(user).to be_invalid
    end

    it 'is valid with matching password confirmation' do
      user.password_confirmation = 'passw0rd'

      expect(user).to be_valid
    end

    it 'is invalid with non matching password confirmation' do
      user.password_confirmation = 'wrongPassword'

      expect(user).to be_invalid
    end
  end
end
