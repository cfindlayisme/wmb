FROM golang:1.22.2 AS builder

WORKDIR /app

COPY . ./
RUN go mod download

RUN apt-get update && apt-get install -y ca-certificates

# Enable CGO for sqllite dependency
RUN CGO_ENABLED=1 GOOS=linux go build -o /application

FROM debian:bookworm-slim

# Copy the C library from the builder image
COPY --from=builder /lib/x86_64-linux-gnu/libc.so.6 /lib/x86_64-linux-gnu/
COPY --from=builder /lib64/ld-linux-x86-64.so.2 /lib64/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /application /application

# Run
CMD ["/application"]