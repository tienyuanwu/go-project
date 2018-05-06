all:
	go get github.com/gin-contrib/cors
	go get github.com/gin-gonic/gin
	go build -o server 

clean:
	go clean
