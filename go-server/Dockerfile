FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/tspo ./cmd

FROM alpine:3.20
LABEL org.opencontainers.image.source=https://github.com/daniel3579/tspo
WORKDIR /app
COPY --from=builder /app/bin/tspo /app/tspo
CMD ["/app/tspo"]