require 'sinatra'
require 'json'

before do
    content_type 'application/json'
end

get '/signin' do
    {
        access_token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJwYW1pdCIsIm5hbWUiOiJQYXlhbSBNIiwiaWF0IjoxNjg5NTgyMjc1fQ.8QlhZQ8pYnY6eSiJ4OZfC7OOhIPJfMUjyUOtqoB6KXE',
        refresh_token: '1337824e-c90b-41d1-a03c-ec6979eb387e',
        token_type: 'Bearer',
        expires_in: 3600
    }.to_json
end
