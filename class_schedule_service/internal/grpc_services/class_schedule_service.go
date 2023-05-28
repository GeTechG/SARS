package grpc_services

import (
	"context"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/services"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
	"github.com/golang/protobuf/ptypes/empty"
)

type ClassScheduleService struct {
	class_schedule_service.UnimplementedClassScheduleServiceServer
}

func (ClassScheduleService) AddClasses(request *class_schedule_service.AddClassesRequest, stream class_schedule_service.ClassScheduleService_AddClassesServer) error {
	for _, class := range request.Classes {
		err := services.AddClass(class, request.GetReplace())
		if err != nil {
			s := err.Error()
			err = stream.Send(&class_schedule_service.AddClassResponse{
				Error:   true,
				Message: &s,
			})
			return err
		} else {
			err = stream.Send(&class_schedule_service.AddClassResponse{
				Error:   false,
				Message: nil,
			})
		}
	}

	return nil
}

func (ClassScheduleService) GetClasses(ctx context.Context, empty *empty.Empty) (*class_schedule_service.GetClassesResponse, error) {
	classes, err := services.GetClasses()
	if err != nil {
		return nil, err
	}
	return &class_schedule_service.GetClassesResponse{Classes: classes}, nil
}

func (ClassScheduleService) GetClass(ctx context.Context, request *class_schedule_service.GetClassRequest) (*class_schedule_service.Class, error) {
	class, err := services.GetClass(request.GetClassId())
	if err != nil {
		return nil, err
	}
	return class, nil
}
