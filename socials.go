package guildedgo

import "errors"

type SocialResponse struct {
	SocialLink `json:"socialLink"`
}

type SocialLink struct {
	Handle    string `json:"handle,omitempty"`
	ServiceID string `json:"serviceId,omitempty"`
	Type      string `json:"type"`
}

type SocialsService interface {
	GetMemberSocialLinks(serverID string, userID string, socialType string) (*SocialLink, error)
}

type socialsEndpoint struct{}

func (e *socialsEndpoint) Default(serverID string, userID string, socialType string) string {
	return guildedApi + "/servers/" + serverID + "/members/" + userID + "/social-links/" + socialType
}

type socialsService struct {
	client    *Client
	endpoints *socialsEndpoint
}

func (s *socialsService) GetMemberSocialLinks(serverID string, userID string, socialType string) (*SocialLink, error) {
	var social SocialResponse
	err := s.client.GetRequestV2(s.endpoints.Default(serverID, userID, socialType), &social)
	if err != nil {
		return nil, errors.New("failed to get social links: " + err.Error())
	}
	return &social.SocialLink, nil
}
