FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init http-server-desafio-korp && \
    go get github.com/prometheus/client_golang/prometheus/promhttp && \
    CGO_ENABLE=0 GOOS=linux go build -o server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/serve .
EXPOSE 8080
CMD ["./server"]