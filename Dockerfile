FROM golang:1.22.4 AS builder

WORKDIR /app

COPY . ./
RUN go mod download

RUN apt-get update && apt-get install -y ca-certificates

# Enable CGO for sqllite dependency
RUN CGO_ENABLED=1 GOOS=linux go build -o /application

FROM debian:12.6-slim

RUN apt-get update && apt-get install -y libc6 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /application /application

# Run
CMD ["/application"]