export PATH := $(PATH):/usr/local/go/bin:$(HOME)/go/bin

proto_compile:
	cd ../service_utilities && make proto_compile

build: proto_compile
	go build cmd/app/main.go

run: proto_compile build
	./main