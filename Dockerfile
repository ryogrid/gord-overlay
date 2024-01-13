FROM golang:1.21.2 as gordol-build

WORKDIR /go/src/app
COPY go.mod /go/src/app
COPY go.sum /go/src/app
COPY . /go/src/app
RUN mv gossip-overlay ../
RUN go mod download
RUN make build

FROM gcr.io/distroless/cc-debian12

COPY --from=gordol-build /go/src/app/gordolctl /
ENTRYPOINT ["/gordolctl"]
CMD ["-l", "", "-n", ""]