package routes

import (
	"context"
	"git.it-college.ru/i21s617/SARS/rest_service/internal/clients"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/class_schedule_service"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/time"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"net/http"
	time2 "time"
)

type ClassRequest struct {
	Date         time.Date `json:"date" binding:"required"`
	Order        int32     `json:"order" binding:"required"`
	Subject      int32     `json:"subject" binding:"required"`
	Teacher      string    `json:"teacher" binding:"required"`
	Group        string    `json:"group"`
	ClassSubject *string   `json:"class_subject,omitempty" binding:"-"`
}

type AddClassRequest struct {
	Classes []*ClassRequest `json:"classes" binding:"required"`
	Replace *bool           `json:"replace,omitempty" binding:"-"`
}

type ClassResponse struct {
	ClassRequest
	Id int64 `json:"id"`
}

func AddClassRequestValidation(ctx context.Context, sl validator.StructLevel) {
	addClassRequest, ok := sl.Current().Interface().(AddClassRequest)
	if ok {
		for _, class := range addClassRequest.Classes {
			exists, err := clients.GetGroupServiceClient().IsGroupExists(ctx, &ldap_service.IsGroupExistsRequest{Group: class.Group})
			if err != nil {
				sl.ReportError(class.Group, "Group", "group", err.Error(), "")
				continue
			}
			if !exists.GetExists() {
				sl.ReportError(class.Group, "Group", "group", "Invalid group name", "")
			}
		}
	}
}

type AddClassesResponse struct {
	Error   bool    `json:"error"`
	Message *string `json:"message,omitempty"`
}

func AddClasses(c *gin.Context) {
	var classesRequest AddClassRequest
	err := c.ShouldBindWith(&classesRequest, binding.JSON)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			var fieldErrors = make([]string, 0, len(errors))
			for _, fieldError := range errors {
				fieldErrors = append(fieldErrors, fieldError.Error())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors:": fieldErrors})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		}
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func GetClasses(c *gin.Context) {
	grpcResponse, err := clients.GetClassScheduleClient().GetClasses(c, &empty.Empty{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	var response = make([]*ClassResponse, 0, len(grpcResponse.Classes))

	for _, c := range grpcResponse.Classes {
		response = append(response, &ClassResponse{
			ClassRequest: ClassRequest{
				Date:         time.Date(c.GetDate().AsTime()),
				Order:        c.GetOrder(),
				Subject:      c.GetSubject(),
				Teacher:      c.GetTeacher(),
				Group:        c.GetGroup(),
				ClassSubject: c.ClassSubject,
			},
			Id: c.GetId(),
		})
	}

	c.JSON(http.StatusOK, response)
}
