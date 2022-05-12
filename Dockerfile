# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as source

WORKDIR $GOPATH/src/
COPY . .

RUN go build -o api ./cmd/main.go

EXPOSE 8080

CMD [ "./api" ]
