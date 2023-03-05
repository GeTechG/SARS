export PATH := $(PATH):/usr/local/go/bin:$(HOME)/go/bin

PACKAGE := git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto

proto_compile:
	python3 ../protobuf_command_gen/gen.py --proto_path ../protobufs \
	--out_path ../service_utilities/pkg/proto \
	--go_package $(PACKAGE) \
	ldap_service/ldap.proto structs/user.proto