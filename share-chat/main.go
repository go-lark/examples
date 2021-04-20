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
	mb := lark.NewMsgBuffer(lark.MsgShareCard)
	msg := mb.BindEmail(email).ShareChat(groupOpenChatID).Build()
	bot.PostMessage(msg)
}
