go mod vendor
docker build -t go-con-client-2 .
kubectl delete deployment gocon-client-2
kubectl delete service gocon-client-2
kubectl apply -f deployment.yml
kubectl apply -f service.yml