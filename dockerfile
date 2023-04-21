FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o faestebytta-bot cmd/faestebytta-bot/main.go

CMD ["./faestebytta-bot"]

