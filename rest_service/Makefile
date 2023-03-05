export PATH := $(PATH):/usr/local/go/bin:/home/i21s617/go/bin

proto_compile:
	cd ../service_utilities && make proto_compile

build: proto_compile
	go build -tags=jsoniter ./cmd/app/main.go

run: proto_compile build
	./main
