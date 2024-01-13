FROM golang:1.21.2 as gord-build

WORKDIR /go/src/app
COPY go.mod /go/src/app
COPY go.sum /go/src/app
RUN go mod download

COPY . /go/src/app
RUN make build

FROM gcr.io/distroless/cc-debian12

COPY --from=gord-build /go/src/app/gordolctl /
ENTRYPOINT ["/gordolctl"]
CMD ["-l", "", "-n", ""]