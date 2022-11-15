FROM golang:1.19-alpine3.15 AS build-env
WORKDIR /src

#ADD . /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

#RUN go clean --modcache
RUN go build -o /app cmd/main.go

FROM alpine:3.15
COPY --from=build-env app /
ENTRYPOINT ["/app"]