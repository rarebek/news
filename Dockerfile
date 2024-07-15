# Step 1: Modules caching
FROM golang:1.23rc1-alpine3.20 AS modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.23rc1-alpine3.20 AS builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app_news
WORKDIR /app_news
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app_news ./cmd/app

# Step 3: Final
FROM scratch
COPY --from=builder /app_news/config /config
COPY --from=builder /app_news/migrations /migrations
COPY --from=builder /bin/app_news /app_news
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app_news"]
