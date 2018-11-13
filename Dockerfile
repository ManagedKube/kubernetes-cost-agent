FROM golang:1.11.0-stretch as builder

COPY . $GOPATH/src/managedkube/agent/
WORKDIR $GOPATH/src/managedkube/agent/

# get dependancies
RUN go get -d -v

# build
RUN go build -o /go/bin/agent

ENTRYPOINT ["/go/bin/agent"]

# # start from scratch
# FROM scratch
#
# # Copy our static executable
# COPY --from=builder /go/bin/mk-agent /mk-agent
# RUN chmod 777 /mk-agent
# ENTRYPOINT ["/mk-agent"]
