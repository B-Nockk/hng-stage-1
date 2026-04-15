FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
