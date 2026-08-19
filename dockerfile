FROM golang:1.26.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /app/main .


FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache \
    ca-certificates \
    wget

RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /app/main .

RUN chown app:app /app/main

USER app

EXPOSE 8080

CMD ["./main"]