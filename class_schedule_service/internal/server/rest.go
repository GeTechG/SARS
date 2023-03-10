package server

import (
	"git.it-college.ru/i21s617/SARS/class_schedule_service/internal/routes"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/sessions"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func RunRestServer() (*http.Server, error) {
	log := logger.GetSugarLogger()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	err := logger.ConnectToGin(router)
	if err != nil {
		return nil, err
	}

	router.Use(cors.Default())

	router.POST("/add_classes", sessions.AuthMiddleware(routes.AddClasses))

	addr := os.Getenv("HOST") + ":" + os.Getenv("REST_PORT")
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
