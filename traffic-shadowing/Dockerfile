FROM golang:1.15.8-alpine3.13 as base

COPY . /app
WORKDIR /app

RUN go build main.go

FROM alpine:3.13

COPY --from=base /app/main /app/main

EXPOSE 8000

CMD ["/app/main"]