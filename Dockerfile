FROM golang:latest

# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY ./go/go.mod ./
COPY ./go/go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-kraken

EXPOSE 8080 

CMD [ "/docker-kraken" ]