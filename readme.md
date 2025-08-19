# CNPG Test

This project is a small playground for CloudNativePG (CNPG).
It spins up a Postgres cluster in minikube, configures TLS/mTLS with cert-manager, and connects to it from a simple Go app using sqlx.

The goal: experiment with CNPG and see how mTLS works in CNPG.

## Setup

1. Start Minikube
   minikube start

2. Install cert-manager
   kubectl apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml

3. Install CNPG Operator
   kubectl apply --server-side -f \
    https://raw.githubusercontent.com/cloudnative-pg/cloudnative-pg/release-1.26/releases/cnpg-1.26.1.yaml

4. Choose your cluster

Simple (password auth):

kubectl apply -f basic-cnpg-cluster.yaml

Advanced (mTLS):

kubectl apply -f mTLS-cnpg-cluster.yaml

5. Build and load the Go app
   podman build -t cnpg_test:latest .
   minikube image load cnpg_test

6. Deploy the app
   kubectl apply -f go-deployment.yaml

7. Test it

Run a tunnel:

minikube tunnel

Check the app:

curl localhost:8080/healtz
curl localhost:8080/time
