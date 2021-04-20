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
	// Upload to Lark at first
	resp, err := bot.UploadImage("/path/to/your.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Send with ImageKey
	mb := lark.NewMsgBuffer(lark.MsgImage)
	msg := mb.BindEmail("xxxxxx@example.com").Image(resp.Data.ImageKey).Build()
	bot.PostMessage(msg)
}
