package guildedgo

import "errors"

type XPObject struct {
	Amount int `json:"amount"`
}

type AwardXPResponse struct {
	Total int `json:"total"`
}

type xpEndpoints struct{}

func (e *xpEndpoints) Default(serverID, userID string) string {
	return guildedApi + "/servers/" + serverID + "/members/" + userID + "/xp"
}

func (e *xpEndpoints) Role(serverID, roleID string) string {
	return guildedApi + "/servers/" + serverID + "/roles/" + roleID + "/xp"
}

type ServerXPService interface {
	AwardXP(serverID, userID string, xpObject *XPObject) (*AwardXPResponse, error)
	SetMemberXP(serverID, userID string, xpObject *XPObject) (*AwardXPResponse, error)
	AwardRoleXP(serverID, roleID string, xpObject *XPObject) error
}

type serverXPService struct {
	endpoints xpEndpoints
	client    *Client
}

var _ ServerXPService = &serverXPService{
	endpoints: xpEndpoints{},
}

func (service *serverXPService) AwardXP(serverID, userID string, xpObject *XPObject) (*AwardXPResponse, error) {
	var response AwardXPResponse
	err := service.client.PostRequestV2(service.endpoints.Default(serverID, userID), &xpObject, &response)
	if err != nil {
		return nil, errors.New("error awarding xp: " + err.Error())
	}

	return &response, nil
}

func (service *serverXPService) SetMemberXP(serverID, userID string, xpObject *XPObject) (*AwardXPResponse, error) {
	var response AwardXPResponse
	err := service.client.PutRequestV2(service.endpoints.Default(serverID, userID), &xpObject, &response)
	if err != nil {
		return nil, errors.New("error setting member xp: " + err.Error())
	}

	return &response, nil
}

func (service *serverXPService) AwardRoleXP(serverID, roleID string, xpObject *XPObject) error {
	err := service.client.PostRequestV2(service.endpoints.Role(serverID, roleID), &xpObject, nil)
	if err != nil {
		return errors.New("error awarding role xp: " + err.Error())
	}

	return nil
}
