package adaptor

import (
	"github.com/nlopes/slack"
)

type api struct {
	bot   *slack.Client
	oauth *slack.Client
}

type API interface {
	PostMessage(channel string, text string) error
	PostMessageWithOptions(change string, text string, msgOption slack.MsgOption) error
	GetUsersFromChannel(channelID string) ([]string, error)
	GetUserGroups() ([]slack.UserGroup, error)
	GetUsersFromUserGroup(groupID string) ([]string, error)
}

func NewAPI(botToken string, oauthToken string) API {
	return &api{slack.New(botToken), slack.New(oauthToken)}
}

func (a *api) PostMessage(channel string, text string) error {
	_, _, err := a.bot.PostMessage(channel, slack.MsgOptionText(text, false))
	return err
}

func (a *api) PostMessageWithOptions(channel string, text string, msgOption slack.MsgOption) error {
	_, _, err := a.bot.PostMessage(channel, slack.MsgOptionText(text, false), msgOption)
	return err
}

func (a *api) GetUsersFromChannel(channelID string) ([]string, error) {
	params := slack.GetUsersInConversationParameters{ChannelID: channelID}
	userIDs, _, err := a.bot.GetUsersInConversation(&params)
	return userIDs, err
}

func (a *api) GetUserGroups() ([]slack.UserGroup, error) {
	groups, err := a.oauth.GetUserGroups()
	return groups, err
}

func (a *api) GetUsersFromUserGroup(groupID string) ([]string, error) {
	userIDs, err := a.oauth.GetUserGroupMembers(groupID)
	return userIDs, err
}
