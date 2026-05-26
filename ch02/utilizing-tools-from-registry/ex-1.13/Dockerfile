FROM golang:1.16

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

RUN go test ./...

ENV PORT=8080

EXPOSE 8080

ENV REQUEST_ORIGIN="https://example.com"

CMD ["./server"]

# docker build -f Dockerfile -t exampl-backend .
# docker run -p 5002:8080 exampl-backend:latest
