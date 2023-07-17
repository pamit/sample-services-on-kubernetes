# Sample Services on Kubernetes

This repository contains 2 sample services, a Ruby authentication service and a Golang order service which calls the authentication service.

You can read the main article on Medium: https://pamit.medium.com/

# Getting started

You can run the apps using Docker Compose:

```
docker compose up

curl -XGET http://localhost:8080/order
{"message":"Order placed for user Payam M (username: pamit) | on 2023-07-17 08:24:35 +0000 UTC"}
```

## Kubernetes (K8s)

When running the apps on a K8s cluster, you need set this environment variable:

```
AUTH_SERVICE_URL=http://ruby-authentication-service:4567

kubectl create configmap currency-conversion --from-literal=AUTH_SERVICE_URL=http://ruby-authentication-service:4567
```
