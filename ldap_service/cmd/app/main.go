package main

import (
	"context"
	"git.it-college.ru/i21s617/SARS/ldap_service/internal/ldap_client"
	"git.it-college.ru/i21s617/SARS/ldap_service/internal/server"
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

	log = logger.GetSugarLogger()

	err = ldap_client.Init()
	if err != nil {
		log.Panic(err)
	}
	defer ldap_client.Shutdown()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	grpcServer, err := server.RunServer()
	if err != nil {
		log.Panic(err)
	}

	log.Infof("Server started at http://0.0.0.0:%s", os.Getenv("PORT"))
	<-done

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	grpcServer.GracefulStop()
	log.Info("Server Exited Properly")
}
