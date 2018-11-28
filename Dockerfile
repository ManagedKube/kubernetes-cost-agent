FROM golang:1.11.0-stretch as builder

COPY . $GOPATH/src/managedkube.com/kubernetes-cost-agent/
WORKDIR $GOPATH/src/managedkube.com/kubernetes-cost-agent/

# get dependancies
RUN go get -d -v

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/agent

# Pricing sheet's folders
RUN mkdir -p /pkg/price
RUN cp -a $GOPATH/src/managedkube.com/kubernetes-cost-agent/pkg/price/prices /pkg/price/

####################
# start from scratch
####################
FROM scratch

# Copy our static executable
COPY --from=builder /go/bin/agent /mk-agent

# Copy pricing sheets
COPY --from=builder /pkg /pkg

ENTRYPOINT ["/mk-agent"]
