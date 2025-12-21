FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o worker main.go

FROM alpine:3.19
COPY --from=builder /app/worker /usr/local/bin/worker
CMD ["worker"]
