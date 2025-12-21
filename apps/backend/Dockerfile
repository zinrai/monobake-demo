FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o backend main.go

FROM alpine:3.19
COPY --from=builder /app/backend /usr/local/bin/backend
EXPOSE 8000
CMD ["backend"]
