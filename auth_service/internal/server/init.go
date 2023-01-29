package server

import (
	"git.it-college.ru/i21s617/SARS/auth_service/internal/routes"
	"git.it-college.ru/i21s617/SARS/auth_service/internal/sessions"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func RunServer() (*http.Server, error) {
	log := logger.GetSugarLogger()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	err := logger.ConnectToGin(router)
	if err != nil {
		return nil, err
	}

	router.Use(cors.Default())

	router.GET("/get_user/:uid", routes.GetUser)
	router.POST("/auth", routes.Auth)
	router.GET("/is_auth", routes.IsAuth)

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
