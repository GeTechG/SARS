package grpc_services

import (
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/services"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
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
