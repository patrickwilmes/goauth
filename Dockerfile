FROM docker.io/golang:alpine AS builder

RUN apk update && \
    apk add --no-cache gcc libc-dev

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 go build -ldflags="-s -w -extldflags '-static'" -tags netgo -installsuffix netgo  -o goauth /app/cmd/goauth

FROM scratch

WORKDIR /app

# This line is necessary for email notification for registering a new user
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/goauth .
COPY --from=builder /app/config/ config/
COPY --from=builder /app/templates/ templates/

EXPOSE 8080

CMD ["/app/goauth"]