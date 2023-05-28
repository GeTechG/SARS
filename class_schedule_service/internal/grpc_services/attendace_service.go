package grpc_services

import (
	"context"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/services"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AttendanceService struct {
	class_schedule_service.UnimplementedAttendanceServiceServer
}

func (AttendanceService) GetAttendance(ctx context.Context, request *class_schedule_service.GetAttendanceRequest) (*class_schedule_service.GetAttendanceResponse, error) {
	attendances, err := services.GetAttendances(request.GetClassId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &class_schedule_service.GetAttendanceResponse{Attendances: attendances}, nil
}

func (AttendanceService) SetAttendance(ctx context.Context, request *class_schedule_service.SetAttendanceRequest) (*class_schedule_service.SetAttendanceResponse, error) {
	var errors = make([]string, 0, len(request.Attendances))
	for _, attendance := range request.GetAttendances() {
		err := services.SetAttendance(attendance)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	return &class_schedule_service.SetAttendanceResponse{Errors: errors}, nil
}
