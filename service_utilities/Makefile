export PATH := $(PATH):/usr/local/go/bin:$(HOME)/go/bin

PACKAGE := git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto

proto_compile:
	rm -rf pkg/proto/*
	python3 ../protobuf_command_gen/gen.py --proto_path ../protobufs \
	--out_path ../service_utilities/pkg/proto \
	--go_package $(PACKAGE) \
	ldap_service/user.proto ldap_service/group.proto \
	structs/user.proto \
	class_schedule_service/class.proto class_schedule_service/attendance.proto class_schedule_service/service.proto