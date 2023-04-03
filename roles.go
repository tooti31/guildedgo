package guildedgo

import (
	"log"
)

type RoleService interface {
	AddMemberToGroup(groupId string, userId string)
	RemoveMemberFromGroup(groupId string, userId string)
}

type roleEndpoints struct{}

func (e *roleEndpoints) GroupMember(groupId, userId string) string {
	return guildedApi + "/groups/" + groupId + "/members/" + userId
}

type roleService struct {
	client    *Client
	endpoints *roleEndpoints
}

var _ RoleService = &roleService{}

func (rs *roleService) AddMemberToGroup(groupId string, userId string) {
	endpoint := rs.endpoints.GroupMember(groupId, userId)

	_, err := rs.client.PutRequest(endpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (rs *roleService) RemoveMemberFromGroup(groupId string, userId string) {
	endpoint := rs.endpoints.GroupMember(groupId, userId)

	_, err := rs.client.DeleteRequest(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
}
