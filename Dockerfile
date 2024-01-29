FROM golang:1.21.2 as gordol-build

WORKDIR /go/src/app
COPY go.mod /go/src/app
COPY go.sum /go/src/app
COPY start_kvs_and_proxy.sh /go/src/app
COPY . /go/src/app/
RUN chmod +x /go/src/app/start_kvs_and_proxy.sh
RUN go mod download
RUN make build

WORKDIR /go/src/app/third/gossip-port-forward
RUN go mod download
RUN go build -o gossip-port-forward gossip-port-forward.go

FROM gcr.io/distroless/cc-debian12

COPY --from=gordol-build /go/src/app/gordolctl /
COPY --from=gordol-build /go/src/app/third/gossip-port-forward/gossip-port-forward /
COPY --from=gordol-build /go/src/app/start_kvs_and_proxy.sh /
WORKDIR /
#ENTRYPOINT ["/gordolctl"]
ENTRYPOINT ["/start_kvs_and_proxy.sh"]
CMD [""]
