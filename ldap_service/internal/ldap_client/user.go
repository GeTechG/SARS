package ldap_client

import (
	"fmt"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/logger"
	"github.com/go-ldap/ldap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *Service) GetUsers(uids []string) ([]*ldap.Entry, error) {
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

	var filterUids = make([]string, 0, len(uids))
	for _, uid := range uids {
		filterUids = append(filterUids, fmt.Sprintf("(uid=%s)", uid))
	}

	request := ldap.NewSearchRequest(
		"dc=it-college,dc=ru",
		ldap.ScopeWholeSubtree,
		ldap.DerefInSearching|ldap.DerefFindingBaseObj,
		0,
		0,
		false,
		fmt.Sprintf("(|%s)", strings.Join(filterUids, "")),
		[]string{},
		nil,
	)

	result, err := s.connect.Search(request)
	if err != nil {
		return nil, err
	}

	if len(result.Entries) == 0 {
		return nil, status.Error(codes.NotFound, "Not found received uids")
	}

	return result.Entries, nil
}

func (s *Service) Auth(uid string, password string) (*ldap.Entry, error) {
	users, err := s.GetUsers([]string{uid})
	if err != nil {
		return nil, err
	}

	user := users[0]
	err = s.connect.Bind(user.DN, password)
	if err != nil {
		return nil, err
	}

	return user, err
}
