# Golang Order Service

This is a simple Golang app which calls the Ruby Authentication Service on port 4567.

You can build the Docker image and push it to your own image respository such as Docker Hub or AWS ECR:

```
# Build image and run a container
docker build -t golang-order-service .
docker run golang-order-service

# Push image to repository
docker tag golang-order-service DOCKER_HUB_USERNAME/golang-order-service:latest
docker push DOCKER_HUB_USERNAME/golang-order-service:latest

Order placed for user Payam M (username: pamit) | on 2023-07-17 18:24:35 +1000 AEST
```

The image is currently hosted on:
https://hub.docker.com/repository/docker/pamitedu/golang-order-service/general
