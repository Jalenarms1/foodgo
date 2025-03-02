FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache \
    pnpm \
    ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd ./client && pnpm install && pnpm run build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/foodgo ./cmd/foodgo
RUN apk add --no-cache ca-certificates

FROM scratch
WORKDIR /root

COPY --from=builder /app/bin .
COPY --from=builder /app/client /root/client

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


EXPOSE 8080

CMD ["./foodgo"]
