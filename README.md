# deploy-nginx
deploy nginx will deploy surprisingly four replicas of nginx on eks cluster with k8s deployment distributed on three nodes

## Setup
These enviroment variables are required:

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

```
$ terraform init
```

## Deploy 
```
$ terraform apply
```

## Output
```cluster_endpoint```: K8s cluster endpoint

```lb_ip```: Load balancer ip for k8s cluster

## Destroy
```
$ terraform destroy
```
