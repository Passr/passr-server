class SessionsController < ApplicationController
  def new
  end

  def create
    user = login(params[:email], params[:password])

    redirect_to dashboard_path
  end

  def destroy
    logout

    redirect_to sign_in_path
  end
end
