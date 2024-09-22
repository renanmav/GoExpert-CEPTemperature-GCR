FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]