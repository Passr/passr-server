Rails.application.routes.draw do
  get '/sign_in', to: 'sessions#new'
  resource :sessions, only: [:create, :destroy]

  resource :dashboard, only: [:show]

  # get '/sign_up', to: 'users#new'

  root to: 'root#index'
end
