https://cloudnative-pg.io/documentation/1.26/quickstart/

minikube start

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml

kubectl apply --server-side -f \
 https://raw.githubusercontent.com/cloudnative-pg/cloudnative-pg/release-1.26/releases/cnpg-1.26.1.yaml

kubectl get secret server-ca-key-pair -o jsonpath="{.data.ca\.crt}" | base64 -d > ca.crt

psql "host=localhost port=5432 dbname=app user=app sslmode=verify-ca sslrootcert=ca.crt"
