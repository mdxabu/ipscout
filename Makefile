BINARY_NAME := ipscout

build:
	go build -o $(BINARY_NAME) main.go

install:
	go install
