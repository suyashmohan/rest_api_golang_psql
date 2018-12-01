.PHONY: default dep build run clean

BINNAME := rest

export GOPATH=$(shell pwd)/vendor
# export GOROOT=/usr/lib/go

default: run

dep:
	go get github.com/julienschmidt/httprouter
	go get github.com/lib/pq
	go get golang.org/x/crypto/bcrypt
	go get github.com/dgrijalva/jwt-go
	go get github.com/davecgh/go-spew/spew

build:
	go build -o ${BINNAME} src/*.go

run:
	go run src/*.go

clean:
	rm -f ${BINNAME}

