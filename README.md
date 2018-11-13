Kube Cost Agent
================

cd ~/go/src/managedkube.com/kube-cost-agent

docker run -it -v ${PWD}:/go/src -v /home/g44/Downloads/gcp-kubeconfig:/root/.kube/config -v ~/Downloads:/opt/Downloads golang:1.11.0-stretch bash

'''
cd ~/go/src/managedkube.com/kube-cost-agent

#export GOROOT=/home/g44/go
#export GOPATH=/home/g44/Documents/managed-kubernetes/kubernetes-cost-attribution/golang/src

go mod init
go run main.go --kubeconfig ~/.kube/config
go build .
'''

## Metrics

```
curl http://localhost:9101/metrics
```

# Docker build

```
docker build gcr.io/managedkube/kubernetes-cost-attribution/agent:dev

docker push gcr.io/managedkube/kubernetes-cost-attribution/agent:dev
```
