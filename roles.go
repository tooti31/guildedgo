package guildedgo

import (
	"log"

	"github.com/itschip/guildedgo/internal/endpoints"
)

type RoleService interface {
	AddMemberToGroup(groupId string, userId string)
	RemoveMemberFromGroup(groupId string, userId string)
}

type roleService struct {
	client *Client
}

var _ RoleService = &roleService{}

func (rs *roleService) AddMemberToGroup(groupId string, userId string) {
	endpoint := endpoints.GroupMemberEndpoint(groupId, userId)

	_, err := rs.client.PutRequest(endpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (rs *roleService) RemoveMemberFromGroup(groupId string, userId string) {
	endpoint := endpoints.GroupMemberEndpoint(groupId, userId)

	_, err := rs.client.DeleteRequest(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
}
