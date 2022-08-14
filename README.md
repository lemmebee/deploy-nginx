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
cd test && go test -v -timeout 2000s infra_test.go
```

## Deploy 
```
$ terraform apply
```

## Destroy
```
$ terraform destroy
```

## Github Workflow

On every push for any branch will trigger ```deploy-terraform``` it will test, deploy then destroy (Will be locked only for main) \
On main and develop pull requests will trigger ```validate-terraform``` it will format, validate and plan

Environment variables are hooked up as secret like the following:
```
env:
  AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
  AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
```