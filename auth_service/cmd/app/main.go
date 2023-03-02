package main

import (
	"context"
	"git.it-college.ru/i21s617/SARS/auth_service/internal/grpc"
	"git.it-college.ru/i21s617/SARS/auth_service/internal/server"
	"git.it-college.ru/i21s617/SARS/auth_service/internal/sessions"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var log *zap.SugaredLogger

func loadEnv() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	err = godotenv.Load(string(exPath) + `/.env`)
	_ = os.Setenv("EX_PATH", exPath)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	loadEnv()

	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Flush()

	err = sessions.InitSessions()
	if err != nil {
		panic(err)
	}
	defer sessions.Shutdown()

	log = logger.GetSugarLogger()

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	srv, err := server.RunServer()
	if err != nil {
		panic(err)
	}

	userServiceConnect, err := grpc.ConnectLDAPToServer()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = userServiceConnect.Close()
	}()

	log.Infof("Server started at http://%s", os.Getenv("HOST")+":"+os.Getenv("PORT"))
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Info("Shutdown Server ...")
}
