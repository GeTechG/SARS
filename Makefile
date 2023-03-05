build-services:
	cd auth_service && make build
	cd ldap_service && make build

run-services:
	systemctl daemon-reload
	systemctl restart sars_auth_service.service
	systemctl restart sars_ldap_service.service
