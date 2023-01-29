package ldap

import (
	"fmt"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"github.com/go-ldap/ldap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetUser(uid string) (*ldap.Entry, error) {
	log := logger.GetSugarLogger()

	errChan := make(chan error)
	s.requestWrapper(func() {
		err := s.connect.Bind(s.bindUsername, s.bindPassword)
		if err != nil {
			errChan <- err
		}
		close(errChan)
	})

	if err := <-errChan; err != nil {
		log.Error(err)
	}

	request := ldap.NewSearchRequest(
		"dc=it-college,dc=ru",
		ldap.ScopeWholeSubtree,
		ldap.DerefInSearching|ldap.DerefFindingBaseObj,
		0,
		0,
		false,
		fmt.Sprintf("(uid=%s)", uid),
		[]string{},
		nil,
	)

	result, err := s.connect.Search(request)
	if err != nil {
		return nil, err
	}

	if len(result.Entries) == 0 {
		return nil, status.Error(codes.NotFound, "Not found received uid")
	}

	return result.Entries[0], nil
}

func (s *Service) Auth(uid string, password string) (*ldap.Entry, error) {
	user, err := s.GetUser(uid)
	if err != nil {
		return nil, err
	}

	err = s.connect.Bind(user.DN, password)
	if err != nil {
		return nil, err
	}

	return user, err
}
