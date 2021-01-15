FROM golang:rc-alpine3.12 AS builder

WORKDIR $GOPATH

COPY . .

# The resulting binary will not be linked to any C libraries
ENV CGO_ENABLED=0

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/challenge06 src/challenge06/main.go

FROM scratch

COPY --from=builder /go/bin/challenge06 /go/bin/challenge06

ENTRYPOINT ["/go/bin/challenge06"]
