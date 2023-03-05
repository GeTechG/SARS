package ldap_client

import (
	"context"
	"fmt"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"github.com/go-ldap/ldap"
	"os"
)

type Service struct {
	connect      *ldap.Conn
	bindUsername string
	bindPassword string
}

var service *Service

func connect() (*ldap.Conn, error) {
	host := os.Getenv("LDAP_HOST")
	conn, err := ldap.DialURL(fmt.Sprintf("ldap://%s:389", host))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *Service) requestWrapper(fn func()) {
	ctx, cancel := context.WithTimeout(context.Background(), ldap.DefaultTimeout)

	go func(ctx context.Context) {
		defer cancel()
		fn()
	}(ctx)

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			logger.GetSugarLogger().Panic("connected LDAP is dead")
		}
	}
}

func Init() error {
	var err error

	conn, err := connect()
	if err != nil {
		return err
	}

	service = &Service{
		connect:      conn,
		bindUsername: os.Getenv("LDAP_USERNAME"),
		bindPassword: os.Getenv("LDAP_PASSWORD"),
	}

	return nil
}

func GetService() *Service {
	return service
}

func Shutdown() {
	service.connect.Close()
}
