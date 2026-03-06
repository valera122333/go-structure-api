FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

ENV PATH=$PATH:/root/go/bin

RUN go build -o app ./cmd/app

CMD ["./app"]