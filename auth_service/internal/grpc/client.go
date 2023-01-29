package grpc

import (
	"git.it-college.ru/i21s617/SARS/auth_service/internal/proto/ldap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

var userServiceClient ldap.UserServiceClient

func ConnectLDAPToServer() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv("LDAP_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userServiceClient = ldap.NewUserServiceClient(conn)

	return conn, nil
}

func GetUserService() ldap.UserServiceClient {
	return userServiceClient
}
