package main

import (
	"fmt"

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
	// GetUserIDByEmail to get openID
	userInfoByEmail, _ := bot.GetUserIDByEmail("xxxxxx@example.com")
	userOpenID := userInfoByEmail.OpenID
	// GetUserInfo with openID
	userInfo, _ := bot.GetUserInfo(userOpenID)
	fmt.Println(userInfo.Avatar) // get user avatar
}
