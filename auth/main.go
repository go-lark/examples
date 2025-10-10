// Package main auth example
package main

import (
	"context"
	"fmt"

	"github.com/go-lark/lark/v2"
)

// should set your own appID and appSecret
const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"
)

func main() {
	ctx := context.Background()
	bot := lark.NewChatBot(appID, appSecret)
	// Single-pass authentication, should renew manually
	resp, err := bot.GetTenantAccessTokenInternal(ctx)
	if err == nil {
		fmt.Println(resp.Expire)
	}
	// your awesome codes ...
}
