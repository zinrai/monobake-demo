FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o frontend main.go

FROM alpine:3.19
COPY --from=builder /app/frontend /usr/local/bin/frontend
EXPOSE 8080
CMD ["frontend"]
