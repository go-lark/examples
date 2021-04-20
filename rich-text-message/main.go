package main

import "github.com/go-lark/lark"

// should set your own appID and appSecret
const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"
)

func main() {
	bot := lark.NewChatBot(appID, appSecret)
	_, _ = bot.GetTenantAccessTokenInternal(true)

	msg := lark.NewMsgBuffer(lark.MsgPost)
	postContent := lark.NewPostBuilder().
		Title("post title").
		TextTag("hello, world", 1, true).
		LinkTag("Google", "https://google.com/").
		AtTag("www", groupOpenChatID).
		ImageTag("d9f7d37e-c47c-411b-8ec6-9861132e6986", 300, 300).
		Render()
	om := msg.BindOpenChatID(groupOpenChatID).Post(postContent).Build()
	_, _ = bot.PostMessage(om)
}
