FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init http-server-projeto-korp && \
    go mod tidy && \
    CGO_ENABLE=0 GOOS=linux go build -o server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]