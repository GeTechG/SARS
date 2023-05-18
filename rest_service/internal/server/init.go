package server

import (
	"git.it-college.ru/i21s617/SARS/rest_service/internal/routes"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/sessions"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
)

func RunServer() (*http.Server, error) {
	log := logger.GetSugarLogger()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidationCtx(routes.AddClassRequestValidation, routes.AddClassRequest{})
	}

	err := logger.ConnectToGin(router)
	if err != nil {
		return nil, err
	}

	router.Use(cors.Default())

	authGroup := router.Group("/auth")
	{
		authGroup.GET("/get_user/:uid", routes.GetUser)
		authGroup.POST("/auth", routes.Auth)
		authGroup.GET("/is_auth", routes.IsAuth)
	}

	classScheduleGroup := router.Group("/class_schedule")
	{
		classScheduleGroup.POST("/add_classes", sessions.AuthMiddleware(routes.AddClasses))
		classScheduleGroup.GET("/get_classes", sessions.AuthMiddleware(routes.GetClasses))
	}

	addr := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	srv := &http.Server{
		Addr:    addr,
		Handler: sessions.GetSessions().LoadAndSave(router),
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv, nil
}
