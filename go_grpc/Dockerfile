FROM golang:1.21-alpine

RUN mkdir /app

ADD . /app 

WORKDIR /app

COPY go.mod ./

RUN go mod download 

RUN go build -o main ./server/server.go 

CMD ["/app/main"] 