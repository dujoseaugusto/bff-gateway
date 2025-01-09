FROM golang:1.20-alpine as builder

WORKDIR /app/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o entrypoint ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/entrypoint .

EXPOSE 8080

CMD ["./entrypoint"]