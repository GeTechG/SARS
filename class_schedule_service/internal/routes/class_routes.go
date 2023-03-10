package routes

import (
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/clients"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/time"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"net/http"
	time2 "time"
)

type ClassRequest struct {
	Date         time.Date `json:"date"`
	Order        int32     `json:"order"`
	Subject      uint32    `json:"subject"`
	Teacher      string    `json:"teacher"`
	Group        string    `json:"group"`
	ClassSubject *string   `json:"class_subject"`
}

type AddClassRequest struct {
	Classes []*ClassRequest `json:"classes"`
	Replace *bool           `json:"replace,omitempty"`
}

type AddClassesResponse struct {
	Error   bool    `json:"error"`
	Message *string `json:"message,omitempty"`
}

func AddClasses(c *gin.Context) {
	var classesRequest AddClassRequest
	err := c.ShouldBindJSON(&classesRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	classespb := make([]*class_schedule_service.Class, 0, len(classesRequest.Classes))
	for _, c := range classesRequest.Classes {
		classespb = append(classespb, &class_schedule_service.Class{
			Date:         timestamppb.New(time2.Time(c.Date)),
			Order:        c.Order,
			Subject:      c.Subject,
			Teacher:      c.Teacher,
			Group:        c.Group,
			ClassSubject: c.ClassSubject,
		})
	}

	stream, err := clients.GetClassScheduleClient().AddClasses(c, &class_schedule_service.AddClassesRequest{
		Classes: classespb,
		Replace: classesRequest.Replace,
	})
	if err != nil {
		_ = c.Error(err)
	}
	waitc := make(chan struct{})

	var response = make([]*AddClassesResponse, 0, len(classesRequest.Classes))

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive class: %v", err)
			}
			response = append(response, &AddClassesResponse{
				Error:   in.Error,
				Message: in.Message,
			})
		}
	}()
	<-waitc

	c.JSON(http.StatusOK, response)
}
