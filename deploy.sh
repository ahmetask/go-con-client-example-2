docker build -t go-con-client-2 .
kubectl delete deployment gocon-client-2
kubectl delete service gocon-client-2
kubectl create -f deployment.yml
kubectl create -f service.yml
