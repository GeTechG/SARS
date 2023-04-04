package clients

import (
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

var classScheduleClient class_schedule_service.ClassScheduleServiceClient

func ConnectToClassScheduleServer() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv("CLASS_SCHEDULE_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	classScheduleClient = class_schedule_service.NewClassScheduleServiceClient(conn)

	return conn, nil
}

func GetClassScheduleClient() class_schedule_service.ClassScheduleServiceClient {
	return classScheduleClient
}
