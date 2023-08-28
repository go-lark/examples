// Package main basic message examples
package main

import "github.com/go-lark/lark"

const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"
)

func main() {
	bot := lark.NewChatBot(appID, appSecret)
	bot.GetTenantAccessTokenInternal(true)
	// Personal Chat, use email
	bot.PostText("hello from example with email", lark.WithEmail("xxxxxx@example.com"))
	// Personal Chat, use lark.WithOpenID
	bot.PostText("hello from example with open_id", lark.WithOpenID("ou_xxxxxx"))

	// Group Chat, use lark.WithChatID
	bot.PostText("hello from example with open_chat_id", lark.WithChatID("oc_xxxxxx"))
	// Group Chat @at message
	bot.PostTextMention("hello from PostTextMention example", "ou_08198ccd6a37644b49f4789c92369c80", lark.WithChatID("oc_xxxxxx"))
	// Group Chat @at all
	bot.PostTextMentionAll("hello from PostTextMentionAll example", lark.WithChatID("oc_xxxxxx"))

	// Send Image
	bot.PostImage("d9f7d37e-c47c-411b-8ec6-9861132e6986", lark.WithChatID("oc_xxxxxx"))
	// Send Group Share Chat
	bot.PostShareChat("oc_1c09434c2264eb52dc7667895a7fac6d", lark.WithChatID("oc_xxxxxx"))
}
