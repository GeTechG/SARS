package services

import (
	"context"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/db"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/sqlc"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
)

func GetAttendances(classId int64) ([]*class_schedule_service.Attendance, error) {
	attendances, err := db.GetQueries().GetAttendances(context.Background(), classId)
	if err != nil {
		return nil, err
	}
	var attendancesProto = make([]*class_schedule_service.Attendance, 0, len(attendances))
	for _, attendance := range attendances {
		attendancesProto = append(attendancesProto, &class_schedule_service.Attendance{
			ClassId: attendance.ClassID,
			UserUid: attendance.UserUid,
			Value:   attendance.Value,
		})
	}
	return attendancesProto, nil
}

func SetAttendance(attendance *class_schedule_service.Attendance) error {
	err := db.GetQueries().SetAttendances(context.Background(), sqlc.SetAttendancesParams{
		ClassID: attendance.GetClassId(),
		UserUid: attendance.GetUserUid(),
		Value:   attendance.GetValue(),
	})
	if err != nil {
		return err
	}
	return nil
}
