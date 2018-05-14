all: 
	make debug

install_lib:
	go get github.com/gin-contrib/cors
	go get github.com/gin-gonic/gin
	go get github.com/lib/pq

build:
	go build -o server 

debug:
	export GIN_MODE=debug
	make install_lib
	go build -o server 

release:
	export GIN_MODE=release
	make install_lib
	make build

clean:
	go clean

run:
	./server

install:
	docker build . -t go-gin
	docker run -i -t -p 8080:8080 go-gin

