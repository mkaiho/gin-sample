# gin-sample

## Environment

- [Docker](https://www.docker.com/)
- [minikube](https://minikube.sigs.k8s.io/docs/) (when run at local k8s)
- [go 1.17](https://golang.org/)
- [make](https://www.gnu.org/software/make/)
- Visual Studio Code Plugin
  - [Remote Development](https://github.com/Microsoft/vscode-remote-release) (on host machine)
  - [Go for Visual Studio Code](https://github.com/golang/vscode-go) (on development container)

## Run and Login development container

1. Click lower left icon on Visual studio code
1. Select "Reopen in Container"
1. Click `ctrl` + `p` and select `>Go: Install/Update Tools`

## Build API server

work on development container

```.sh
make build-server
```

## Run API server on local host machine

```.sh
./build/server
```

## Run API server on local k8s cluster with minikube

work on local host machine

### Start minikube

```.sh
minikube start --driver=hyperkit
minikube addons enable ingress
```

### Build and Push docker image

```.sh
docker build --rm -t mkaihou/gin-sample:1.0 .
docker push mkaihou/gin-sample:1.0
```

### Run API server on local minikube cluster

```.sh
kubectl apply -f ./kubernetes/devlopment/namespace.yaml
kubectl apply -f ./kubernetes/devlopment
```

### Add ingress address to local hosts

```.sh
## HOSTS and ADDRESS add to hosts file
kubectl get ingress -n gin-sample-dev
```

### Request to pods

```.sh
curl <ingress host name>
```
