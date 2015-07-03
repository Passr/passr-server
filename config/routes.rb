Rails.application.routes.draw do
  get '/sign_in', to: 'sessions#new'
  post '/sign_up', to: 'users#new'

  resource :sessions, only: [:create, :destroy]
  resource :dashboard, only: [:show]
  resource :users, only: [:create]

  root to: 'root#index'
end
