docker build -t goproject:test .
kubectl delete  -f ~/Downloads/goproject/deployment.yaml
kubectl apply  -f ~/Downloads/goproject/supervisord_config.yaml
kubectl apply  -f ~/Downloads/goproject/deployment.yaml
kubectl get pods
