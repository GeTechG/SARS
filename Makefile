build-services:
	cd rest_service && make build
	cd ldap_service && make build
	cd class_schedule_service && make build

run-services:
	systemctl daemon-reload
	systemctl restart sars_rest_service.service
	systemctl restart sars_ldap_service.service
	systemctl restart sars_class_schedule_service.service
