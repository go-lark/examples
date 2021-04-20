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
	// Single-pass authentication, should renew manually
	resp, err := bot.GetTenantAccessTokenInternal(true)
	if err == nil {
		fmt.Println(resp.Expire)
	}
	// Automatically renew token
	bot.StartHeartbeat()
}
