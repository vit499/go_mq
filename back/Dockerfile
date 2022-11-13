FROM golang:1.19-alpine3.15 AS builder
#RUN apk --no-cache add gcc g++ make git
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN go get ./...
RUN go build -o /main cmd/main.go
#RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go
FROM alpine:3.15

#RUN apk --no-cache add ca-certificates
#WORKDIR /usr/bin
COPY --from=builder main /bin/main
#EXPOSE 80
#ENTRYPOINT /go/bin/web-app --port 80
ENTRYPOINT ["/bin/main"]