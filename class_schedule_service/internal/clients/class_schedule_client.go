package clients

import (
	"fmt"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

var classScheduleClient class_schedule_service.ClassScheduleServiceClient

func ConnectToClassScheduleServer() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", os.Getenv("GRPC_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	classScheduleClient = class_schedule_service.NewClassScheduleServiceClient(conn)

	return conn, nil
}

func GetClassScheduleClient() class_schedule_service.ClassScheduleServiceClient {
	return classScheduleClient
}
