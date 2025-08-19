# CNPG Test

This project is a small playground for CloudNativePG (CNPG).
It spins up a Postgres cluster in minikube, configures TLS/mTLS with cert-manager, and connects to it from a simple Go app using sqlx.

The goal: experiment with CNPG and see how mTLS works in CNPG.

## Setup

1. Start Minikube
   ```bash
   minikube start
   ```
3. Install cert-manager
   ```bash
   kubectl apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml
   ```
5. Install CNPG Operator
   ```bash
   kubectl apply --server-side -f \
    https://raw.githubusercontent.com/cloudnative-pg/cloudnative-pg/release-1.26/releases/cnpg-1.26.1.yaml
   ```
7. Choose your cluster
   - Simple (password auth):
      ```bash
         kubectl apply -f basic-cnpg-cluster.yaml
      ```
   - Advanced (mTLS):
      ```
      kubectl apply -f mTLS-cnpg-cluster.yaml
      ```
5. Build and load the Go app
   ```bash
   podman build -t cnpg_test:latest .
   minikube image load cnpg_test
   ```
7. Deploy the app
   ```bash
   kubectl apply -f go-deployment.yaml
   ```

9. Test it
   - Run a tunnel:
      ```bash
      minikube tunnel
      ```
   - Check the app:
      ```bash
      curl localhost:8080/healtz
      curl localhost:8080/time
      ```
