require 'spec_helper'

describe UsersController, type: :controller do
  describe 'POST #create' do
    let(:request_params) do
      {
        user: {
          email: 'john.doe@passr.io',
          password: 'passw0rd',
          password_confirmation: 'passw0rd'
        }
      }
    end

    it 'creates a new user' do
      expect do
        post :create, request_params
      end.to change(User, :count).by(1)

      expect(User.first).to have_attributes(
        email: 'john.doe@passr.io'
      )

      expect(response).to redirect_to(:dashboard)
      expect(controller.current_user.email).to eq('john.doe@passr.io')
    end
  end
end
