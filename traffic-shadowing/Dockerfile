FROM golang:1.15.8

COPY . /app
WORKDIR /app

RUN go build main.go

EXPOSE 8000

CMD ["/app/main"]