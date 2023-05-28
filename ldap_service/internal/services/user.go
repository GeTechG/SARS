package services

import (
	"context"
	"git.it-college.ru/i21s617/SARS/ldap_service/internal/ldap_client"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/structs"
	"github.com/go-ldap/ldap"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var isStudent = regexp.MustCompile(`i(\d+)s(\d+)`)
var isTeacher = regexp.MustCompile(`t(\d+)`)

func entryToUser(entry *ldap.Entry) *structs.User {
	user := structs.User{
		Uid: entry.GetAttributeValue("uid"),
		Cn:  entry.GetAttributeValue("cn"),
	}

	uid := user.GetUid()
	switch {
	case isStudent.MatchString(uid):
		user.UserType = "student"
	case isTeacher.MatchString(uid):
		user.UserType = "teacher"
	}

	return &user
}

type UserService struct {
	ldap_service.UnimplementedUserServiceServer
}

func (s *UserService) GetUsers(ctx context.Context, request *ldap_service.GetUsersRequest) (*ldap_service.GetUsersResponse, error) {
	ldapUsers, err := ldap_client.GetService().GetUsers(request.GetUids())
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	var users = make([]*structs.User, 0, len(ldapUsers))
	for _, ldapUser := range ldapUsers {
		users = append(users, entryToUser(ldapUser))
	}

	return &ldap_service.GetUsersResponse{
		Users: users,
	}, nil
}

func (s *UserService) Auth(ctx context.Context, request *ldap_service.AuthUserRequest) (*ldap_service.AuthUserResponse, error) {
	entry, err := ldap_client.GetService().Auth(request.GetUid(), request.GetPassword())
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidCredentials) {
			return &ldap_service.AuthUserResponse{
				Valid: false,
				User:  nil,
			}, nil
		}
		return nil, err
	}

	user := entryToUser(entry)

	return &ldap_service.AuthUserResponse{
		Valid: true,
		User:  user,
	}, nil
}
