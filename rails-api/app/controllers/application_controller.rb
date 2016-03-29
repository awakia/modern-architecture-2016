class ApplicationController < ActionController::API
  def index
    render json: { message: "hello rails server!" }
  end
end
