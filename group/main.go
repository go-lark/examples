package main

import (
	"github.com/go-lark/lark"
)

// should set your own appID and appSecret
const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"
)

func main() {
	bot := lark.NewChatBot(appID, appSecret)
	bot.GetTenantAccessTokenInternal(true)
	userInfo, _ := bot.GetUserIDByEmail("xxxxxx@example.com")
	userOpenID := userInfo.OpenID
	// CreateGroup
	groupInfo, _ := bot.CreateGroup("Group Name", "Group Description", []string{userOpenID})
	groupOpenChatID := groupInfo.OpenChatID

	// GetGroupInfo
	bot.GetGroupInfo(groupOpenChatID)

	// AddGroupMember
	bot.AddGroupMember(groupOpenChatID, []string{userOpenID})

	// DeleteGroupMember
	bot.DeleteGroupMember(groupOpenChatID, []string{userOpenID})

	// GetGroupList
	bot.GetGroupList(1, 10) // pageNum，pageSize
}
