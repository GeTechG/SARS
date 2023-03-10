package services

import (
	"context"
	"database/sql"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/db"
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/sqlc"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
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
