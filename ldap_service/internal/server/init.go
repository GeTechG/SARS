package server

import (
	"errors"
	"fmt"
	"git.it-college.ru/i21s617/SARS/ldap_service/internal/services"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"strconv"
)

func RunServer() (*grpc.Server, error) {
	log := logger.GetSugarLogger()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, errors.New("invalid parse port")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := services.UserService{}
	ldap_service.RegisterUserServiceServer(grpcServer, &userService)

	go func() {
		if err := grpcServer.Serve(lis); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return grpcServer, nil
}
