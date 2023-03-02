export PATH := $(PATH):/usr/local/go/bin:$(HOME)/go/bin

PACKAGE := git.it-college.ru/i21s617/SARS/ldap_service/internal/proto

proto_compile:
	python3 ../protobuf_command_gen/gen.py --proto_path ../protobufs \
	--out_path ../ldap_service/internal/proto \
	--go_package $(PACKAGE) \
	ldap/ldap.proto structs/user.proto

build: proto_compile
	go build cmd/app/main.go

run: proto_compile
	go run cmd/app/main.go