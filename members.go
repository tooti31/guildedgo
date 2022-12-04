package guildedgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/itschip/guildedgo/endpoints"
)

type ServerMember struct {
	User User `json:"user"`

	// (must have unique items true)
	RoleIds []int `json:"roleIds"`

	Nickname string `json:"nickname,omitempty"`

	// The ISO 8601 timestamp that the member was created at
	JoinedAt string `json:"joinedAt"`

	// (default false)
	IsOwner bool `json:"isOwner,omitempty"`
}

type ServerMemberSummary struct {
	User UserSummary `json:"user"`

	// (must have unique items true)
	RoleIds []int `json:"roleIds"`
}

type User struct {
	// The ID of the user
	Id string `json:"id"`

	// The type of user. If this property is absent, it can assumed to be of type user
	Type string `json:"type,omitempty"`

	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar,omitempty"`

	// The banner image associated with the user
	Banner string `json:"banner,omitempty"`

	// The ISO 8601 timestamp that the user was created at
	CreatedAt string `json:"createdAt"`
}

type UserSummary struct {
	// The ID of the user
	Id string `json:"id"`

	//  The type of user. If this property is absent, it can assumed to be of type user
	Type string `json:"type,omitempty"`

	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar,omitempty"`
}

type ServerMemberBan struct {
	User UserSummary `json:"user"`

	// The reason for the ban as submitted by the banner
	Reason string `json:"reason,omitempty"`

	// The ID of the user who created this server member ban
	CreatedBy string `json:"createdBy"`

	// The ISO 8601 timestamp that the server member ban was created at
	CreatedAt string `json:"createdAt"`
}

type NicknameResponse struct {
	Nickname string `json:"nickname"`
}

type ServerMemberResponse struct {
	Member ServerMember `json:"member"`
}

type MembersService interface {
	UpdateMemberNickname(userId string, nickname string) (*NicknameResponse, error)
	DeleteMemberNickname(userId string) error
	GetServerMember(serverId string, userId string) (*ServerMember, error)
	KickMember(userId string) error
	BanMember(userId string, reason string) (*ServerMemberBan, error)
	IsMemberBanned(userId string) (*ServerMemberBan, error)
	UnbanMember(userId string) error
	GetServerMembers() (*[]ServerMemberSummary, error)
}

type membersService struct {
	client *Client
}

var _ MembersService = &membersService{}

func (ms *membersService) UpdateMemberNickname(userId string, nickname string) (*NicknameResponse, error) {
	endpoint := endpoints.MemberNicknameEndpoint(ms.client.ServerID, userId)

	body := &NicknameResponse{
		Nickname: nickname,
	}

	resp, err := ms.client.PutRequest(endpoint, body)
	if err != nil {
		return nil, err
	}

	var nick NicknameResponse

	err = json.Unmarshal(resp, &nick)
	if err != nil {
		return nil, err
	}

	return &nick, nil
}

func (ms *membersService) DeleteMemberNickname(userId string) error {
	endpoint := endpoints.MemberNicknameEndpoint(ms.client.ServerID, userId)

	_, err := ms.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to delete member nickname. Error: %s", err.Error()))
	}

	return nil
}

func (ms *membersService) GetServerMember(serverId string, userId string) (*ServerMember, error) {
	endpoint := endpoints.ServerMemberEndpoint(serverId, userId)

	var member ServerMemberResponse
	err := ms.client.GetRequestV2(endpoint, &member)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New(fmt.Sprintf("Failed to get member. Error: %s", err.Error()))
	}

	return &member.Member, nil
}

func (ms *membersService) KickMember(userId string) error {
	endpoint := endpoints.ServerMemberEndpoint(ms.client.ServerID, userId)

	_, err := ms.client.DeleteRequest(endpoint)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to kick member. Error: %s", err.Error()))
	}

	return nil
}

func (ms *membersService) BanMember(userId string, reason string) (*ServerMemberBan, error) {
	endpoint := endpoints.MemberBanEndpoint(ms.client.ServerID, userId)

	// No need to build a struct here
	body := map[string]string{
		"reason": reason,
	}

	var ban ServerMemberBan
	err := ms.client.PostRequestV2(endpoint, body, &ban)
	if err != nil {
		return nil, err
	}

	return &ban, nil
}

func (ms *membersService) IsMemberBanned(userId string) (*ServerMemberBan, error) {
	// Do we want to use the serverID from the config, or manually input it?
	endpoint := endpoints.ServerMemberEndpoint(ms.client.ServerID, userId)

	var ban ServerMemberBan
	err := ms.client.GetRequestV2(endpoint, &ban)
	if err != nil {
		return nil, err
	}

	return &ban, nil
}

func (ms *membersService) UnbanMember(userId string) error {
	endpoint := endpoints.MemberBanEndpoint(ms.client.ServerID, userId)

	_, err := ms.client.DeleteRequest(endpoint)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return nil
}

func (ms *membersService) GetServerMembers() (*[]ServerMemberSummary, error) {
	endpoint := endpoints.MemberEndpoint(ms.client.ServerID)

	var members []ServerMemberSummary
	err := ms.client.GetRequestV2(endpoint, &members)
	if err != nil {
		return nil, err
	}

	return &members, nil
}
