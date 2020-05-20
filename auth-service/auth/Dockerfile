FROM golang:1.14.3-alpine as build

COPY . /app
WORKDIR /app
RUN go build -o auth main.go 
RUN chmod +x ./auth

FROM alpine:latest
LABEL maintainer="hexo-shikanon-blog <shikanon@tensorbytes.com>"

COPY --from=build /app /app
EXPOSE 8000
CMD ["/app/auth"]