FROM golang:1.23

WORKDIR /app

COPY . .

# CMD ["go", "run" "cmd/main.go"]