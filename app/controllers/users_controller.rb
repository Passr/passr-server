class UsersController < ApplicationController
  def new
  end

  def create
    user = User.create(
      params.require(:user).permit(:email, :password, :password_confirmation)
    )

    if user.errors.any?
      flash.now[:error] = "Failed to create new user"
      render :new
    else
      auto_login(user)
      redirect_to dashboard_path
    end
  end
end
