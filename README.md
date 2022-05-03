## schedule-nginx-deployment
this branch needs k8s cluster
```
$ brew install kind
```
```
$ kind create cluster --name nginx --config kind-config.yaml
```
```
$ kind get clusters
```
```
$  kubectl cluster-info --context kind-nginx
```