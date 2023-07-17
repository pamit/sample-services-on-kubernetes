# Ruby Authentication Service

This is a simple Sinatra app which exposes a sample OAuth2 token response, which is consumed by a Golang client.

You can build the Docker image and push it to your own image respository such as Docker Hub or AWS ECR:

```
# Build image and run a container
docker build -t ruby-authentication-service .
docker run -p 4567:4567 ruby-authentication-service

curl -XGET http://127.0.0.1:4567/signin
{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJwYW1pdCIsIm5hbWUiOiJQYXlhbSBNIiwiaWF0IjoxNjg5NTgyMjc1fQ.8QlhZQ8pYnY6eSiJ4OZfC7OOhIPJfMUjyUOtqoB6KXE","refresh_token":"1337824e-c90b-41d1-a03c-ec6979eb387e","token_type":"Bearer","expires_in":3600}

# Push image to repository
docker tag ruby-authentication-service DOCKER_HUB_USERNAME/ruby-authentication-service:latest
docker push DOCKER_HUB_USERNAME/ruby-authentication-service:latest
```

The image is currently hosted on:
https://hub.docker.com/repository/docker/pamitedu/ruby-authentication-service/general
