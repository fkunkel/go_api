# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS builder
WORKDIR /go/src/github.com/fkunkel/go_api/
COPY main.go ./
COPY go.mod ./
RUN echo $(ls -l /go/src/github.com/fkunkel/go_api/)

RUN go mod tidy
RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/src/github.com/fkunkel/go_api/app ./

EXPOSE 8000
CMD ["./app"]
