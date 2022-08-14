# deploy-nginx
deploy nginx will deploy surprisingly four replicas of nginx on eks cluster using helm release distributed on three nodes

## Setup
These enviroment variables are required:

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

```
$ terraform init
```

## Test

```
cd test && go test -v -timeout 3000s infra_test.go
```

## Deploy 
```
$ terraform apply
```

## Destroy
```
$ terraform destroy
```