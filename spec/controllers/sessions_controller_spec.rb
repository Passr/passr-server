require 'rails_helper'

RSpec.describe SessionsController, type: :controller do
  describe 'GET #new' do
    before do
      get :new
    end

    it 'responds successfully' do
      expect(response.status).to eq(200)
    end
  end

  describe 'POST #create' do
    let(:user) { FactoryGirl.create(:user) }

    before do
      post :create, {email: user.email, password: 'password'}
    end

    it 'signs into the users account' do
      expect(controller.current_user).to eq(user)
    end

    it 'redirects to the dashboard' do
      expect(response).to redirect_to(:dashboard)
    end
  end

  describe 'DELETE #destroy' do
    let(:user) { FactoryGirl.create(:user, password: 'password') }

    before do
      login_user(user)
    end

    it 'signs the user out of their account' do
      delete :destroy

      expect(controller.current_user).to be_nil
      expect(response).to redirect_to(:sign_in)
    end
  end
end
