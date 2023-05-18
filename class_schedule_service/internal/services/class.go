package services

import (
	"context"
	"database/sql"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/db"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/sqlc"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AddClass(class *class_schedule_service.Class, replace bool) error {
	if replace {
		err := db.GetQueries().ReplaceClass(context.Background(), sqlc.ReplaceClassParams{
			Date:    class.Date.AsTime(),
			Order:   class.Order,
			Subject: int32(class.Subject),
			Teacher: class.Teacher,
			Group:   class.Group,
			ClassSubject: sql.NullString{
				String: class.GetClassSubject(),
				Valid:  true,
			},
		})
		if err != nil {
			return err
		}
	} else {
		err := db.GetQueries().CreateClass(context.Background(), sqlc.CreateClassParams{
			Date:    class.Date.AsTime(),
			Order:   class.Order,
			Subject: int32(class.Subject),
			Teacher: class.Teacher,
			Group:   class.Group,
			ClassSubject: sql.NullString{
				String: class.GetClassSubject(),
				Valid:  true,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func GetClasses() ([]*class_schedule_service.Class, error) {
	classes, err := db.GetQueries().GetClasses(context.Background())
	if err != nil {
		return nil, err
	}
	classespb := make([]*class_schedule_service.Class, 0, len(classes))
	for _, class := range classes {
		classespb = append(classespb, &class_schedule_service.Class{
			Date:         timestamppb.New(class.Date),
			Order:        class.Order,
			Subject:      class.Subject,
			Teacher:      class.Teacher,
			Group:        class.Group,
			ClassSubject: &class.ClassSubject.String,
		})
	}

	return classespb, nil
}
