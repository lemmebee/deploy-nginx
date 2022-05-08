## deploy-nginx
this needs kind k8s cluster
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
