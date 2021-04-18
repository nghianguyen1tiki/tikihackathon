FROM golang:1.16-buster

WORKDIR /app
ADD . .
RUN go build -o bin/app cmd/app/main.go

CMD ["./bin/app"]
