Kubernetes Cost Agent
================

ManagedKube’s Kubernetes Cost Attribution is an application to help you understand the cost of your Kubernetes clusters, from what each namespace or pod costs to how much disks or network traffic in the cluster costs, so you can manage your budget and optimize your cloud spend.

This is a service that listens to the Kubernetes API server and generates metrics about cost related items that a cluster will incur.

The metrics are exported on the HTTP endpoint /metrics on the listening port (default 9101). They are served as plaintext. They are designed to be consumed either by Prometheus itself or by a scraper that is compatible with scraping a Prometheus client endpoint. You can also open /metrics in a browser to see the raw metrics.


## Building

'''
go mod init
go run main.go --kubeconfig ~/.kube/config
go build .
'''

### Test

```
go test ./...
```

## Metrics

```
curl http://localhost:9101/metrics
```

# Docker build

```
docker build -t gcr.io/managedkube/kubernetes-cost-attribution/agent:dev .

docker push gcr.io/managedkube/kubernetes-cost-attribution/agent:dev
```

# Install Golang

https://golang.org/doc/install#install
