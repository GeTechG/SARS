package server

import (
	"errors"
	"fmt"
	"git.it-college.ru/i21s617/SARS/ldap_service/internal/services"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
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

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_zap.StreamServerInterceptor(logger.GetLogger()),
	)), grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_zap.UnaryServerInterceptor(logger.GetLogger()),
	)))

	userService := services.UserService{}
	groupService := services.GroupService{}
	ldap_service.RegisterUserServiceServer(grpcServer, &userService)
	ldap_service.RegisterGroupServer(grpcServer, &groupService)

	go func() {
		if err := grpcServer.Serve(lis); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return grpcServer, nil
}
