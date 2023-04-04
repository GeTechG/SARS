package services

import (
	"context"
	"git.it-college.ru/i21s617/SARS/ldap_service/internal/ldap_client"
	"git.it-college.ru/i21s617/SARS/service_utilities/pkg/proto/ldap_service"
	"github.com/golang/protobuf/ptypes/empty"
)

type GroupService struct {
	ldap_service.UnimplementedGroupServiceServer
}

func (GroupService) GetListGroups(ctx context.Context, empty *empty.Empty) (*ldap_service.GetListGroupsResponse, error) {
	entries, err := ldap_client.GetService().GetListGroups()
	if err != nil {
		return nil, err
	}

	var groups = make([]string, 0, len(entries))

	for _, entry := range entries {
		groups = append(groups, entry.GetAttributeValue("cn"))
	}

	return &ldap_service.GetListGroupsResponse{Groups: groups}, nil
}

func (GroupService) GetGroupMembers(ctx context.Context, request *ldap_service.GetGroupMembersRequest) (*ldap_service.GetGroupMembersResponse, error) {
	entry, err := ldap_client.GetService().GetGroupMembers(request.GetGroupName())
	if err != nil {
		return nil, err
	}

	members := entry.GetAttributeValues("memberUid")

	return &ldap_service.GetGroupMembersResponse{Members: members}, nil
}

func (GroupService) IsGroupExists(ctx context.Context, request *ldap_service.IsGroupExistsRequest) (*ldap_service.IsGroupExistsResponse, error) {
	exists, err := ldap_client.GetService().IsGroupExists(request.Group)
	if err != nil {
		return nil, err
	}

	return &ldap_service.IsGroupExistsResponse{Exists: exists}, nil
}
