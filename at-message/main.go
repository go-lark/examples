// Package main at message examples
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
	openID := "ou_xxxxxxx" // you may get this from incoming message or other API
	mb := lark.NewMsgBuffer(lark.MsgText)
	// Method 1
	text := fmt.Sprintf("AT Message 1 <at user_id=\"%s\">%s</at>", openID, "name")
	msg1 := mb.BindOpenChatID("oc_xxxxxx").Text(text).Build()
	bot.PostMessage(msg1)
	mb.Clear()
	// Method 2
	tb2 := lark.NewTextBuilder().Text("AT Message 2").Mention(openID)
	msg2 := mb.BindOpenChatID("oc_xxxxxx").Text(tb2.Render()).Build()
	bot.PostMessage(msg2)
	mb.Clear()
	// AT ALL, should use cautiously to prevent disturbing
	tb3 := lark.NewTextBuilder().Text("AT ALL").MentionAll()
	msgAll := mb.BindOpenChatID("oc_xxxxxx").Text(tb3.Render()).Build()
	bot.PostMessage(msgAll)
}
