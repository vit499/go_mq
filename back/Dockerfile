FROM golang:1.19-alpine3.15 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /main cmd/main.go

FROM alpine:3.15
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]