build-k8:
	podman build -t cnpg_test:latest .
	minikube image load docker.io/library/cnpg_test:latest


