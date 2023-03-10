export PATH := $(PATH):/usr/local/go/bin:$(HOME)/go/bin

proto_compile:
	cd ../service_utilities && make proto_compile

build: proto_compile sql_compile
	go build -tags=jsoniter ./cmd/app/main.go

run: proto_compile build
	./main

sql_compile:
	sqlc generate