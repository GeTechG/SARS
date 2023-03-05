package grpc_client

import (
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

var userServiceClient ldap_service.UserServiceClient

func ConnectLDAPToServer() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv("LDAP_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userServiceClient = ldap_service.NewUserServiceClient(conn)

	return conn, nil
}

func GetUserService() ldap_service.UserServiceClient {
	return userServiceClient
}
