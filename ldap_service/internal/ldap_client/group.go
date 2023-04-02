package ldap_client

import (
	"fmt"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"github.com/go-ldap/ldap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetListGroups() ([]*ldap.Entry, error) {
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
		"ou=groups,dc=it-college,dc=ru",
		ldap.ScopeWholeSubtree,
		ldap.DerefInSearching|ldap.DerefFindingBaseObj,
		0,
		0,
		false,
		"(cn=*)",
		[]string{"cn"},
		nil,
	)

	result, err := s.connect.Search(request)
	if err != nil {
		return nil, err
	}

	return result.Entries, nil
}

func (s *Service) GetGroupMembers(group string) (*ldap.Entry, error) {
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
		fmt.Sprintf("cn=%s,ou=groups,dc=it-college,dc=ru", group),
		ldap.ScopeWholeSubtree,
		ldap.DerefInSearching|ldap.DerefFindingBaseObj,
		0,
		0,
		false,
		"(cn=*)",
		[]string{"memberUid"},
		nil,
	)

	result, err := s.connect.Search(request)
	if err != nil {
		return nil, err
	}

	if len(result.Entries) == 0 {
		return nil, status.Error(codes.NotFound, "Not found received group")
	}

	return result.Entries[0], nil
}

func (s *Service) IsGroupExists(group string) (bool, error) {
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
		"ou=groups,dc=it-college,dc=ru",
		ldap.ScopeWholeSubtree,
		ldap.DerefInSearching|ldap.DerefFindingBaseObj,
		0,
		0,
		false,
		fmt.Sprintf("(cn=%s)", group),
		[]string{},
		nil,
	)

	result, err := s.connect.Search(request)
	if err != nil {
		return false, err
	}

	return len(result.Entries) != 0, nil
}
