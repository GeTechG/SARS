package services

import (
	"context"
	ldaps "git.it-college.ru/i21s617/SARS/ldap_service/internal/ldap"
	ldappb "git.it-college.ru/i21s617/SARS/ldap_service/internal/proto/ldap"
	"git.it-college.ru/i21s617/SARS/ldap_service/internal/proto/structs"
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
	ldappb.UnimplementedUserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, request *ldappb.GetUserRequest) (*ldappb.GetUserResponse, error) {
	ldapUser, err := ldaps.GetService().GetUser(request.GetUid())
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	user := entryToUser(ldapUser)

	return &ldappb.GetUserResponse{
		User: user,
	}, nil
}

func (s *UserService) Auth(ctx context.Context, request *ldappb.AuthUserRequest) (*ldappb.AuthUserResponse, error) {
	entry, err := ldaps.GetService().Auth(request.GetUid(), request.GetPassword())
	if err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidCredentials) {
			return &ldappb.AuthUserResponse{
				Valid: false,
				User:  nil,
			}, nil
		}
		return nil, err
	}

	user := entryToUser(entry)

	return &ldappb.AuthUserResponse{
		Valid: true,
		User:  user,
	}, nil
}
