FROM golang:1.20-alpine3.18

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["./main"]