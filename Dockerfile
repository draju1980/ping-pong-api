FROM golang:1.22.1
WORKDIR /app

COPY go.mod ./
RUN go mod tidy
RUN go mod download

COPY *.go ./

RUN go build -o /ping-pong-api

EXPOSE 8080

CMD [ "/ping-pong-api" ]