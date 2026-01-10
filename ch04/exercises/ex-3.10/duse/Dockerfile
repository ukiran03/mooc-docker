FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app
RUN echo "appuser:x:10001:10001:appuser:/:" > /etc_passwd
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /usr/src/app/duse .
RUN chmod +x /usr/src/app/duse

FROM scratch

COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder /usr/src/app/duse /duse
USER appuser

ENTRYPOINT ["/duse"]
