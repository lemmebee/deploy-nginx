# deploy-nginx
deploy nginx will deploy surprisingly four replicas of nginx on kind cluster with k8s deployment

## Setup
These enviroment variables are required:

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

Create local kind cluster:

```
brew install kind
kind create cluster --name nginx --config kind-config.yaml
kind get clusters
kubectl cluster-info --context kind-nginx
```

You need to get cluster configurations from ```~/.kube/config``` and apply it on ```terraform.tfvars``` then you can run:

```
terraform init
```

## Deploy 
```
terraform apply
```

## Destroy
```
terraform destroy
```