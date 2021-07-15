FROM golang:1.15
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o ./bin/ayr cmd/ayr/main.go
CMD ["./bin/ayr"]