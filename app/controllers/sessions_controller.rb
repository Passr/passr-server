class SessionsController < ApplicationController
  def new
  end

  def create
    user = login(params[:email], params[:password])

    if user
      redirect_to dashboard_path
    else
      flash[:error] = "Invalid email and/or password"
      render :new
    end
  end

  def destroy
    logout

    redirect_to sign_in_path
  end
end
