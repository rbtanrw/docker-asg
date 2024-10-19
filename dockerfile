FROM golang:1.22.1-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o go-web main.go

EXPOSE 8080

CMD ["./go-web"]