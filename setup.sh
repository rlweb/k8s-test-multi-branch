minikube start --vm-driver hyperkit
eval $(minikube docker-env) && docker build -t example-1:latest ./example-1
kubectl apply --namespace example-1 -f ./example-1/k8s.yaml
eval $(minikube docker-env) && docker build -t example-2:latest ./example-2
kubectl apply --namespace example-2 -f ./example-2/k8s.yaml
eval $(minikube docker-env) && docker build -t staging-proxy:latest ./staging-proxy
kubectl apply -f ./staging-proxy/k8s.yaml
echo "\n---\nTesting Branch 1\n---\n"
curl $(minikube service staging-proxy --url) --header 'Host: example-1.staging.test.co.uk'
echo "\n---\nTesting Branch 2\n---\n"
curl $(minikube service staging-proxy --url) --header 'Host: example-2.staging.test.co.uk'
echo "\n---\nTesting Branch which cannot be found!\n---\n"
curl $(minikube service staging-proxy --url) --header 'Host: example-4.staging.test.co.uk'
