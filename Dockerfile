FROM golang:latest

RUN mkdir -p /app
WORKDIR /app

ADD . /app
ENV GIN_MODE=release
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get github.com/lib/pq
RUN go build -o server

EXPOSE 8080

CMD ["./server"]
