package main

import (
	"github.com/go-lark/lark"
	"github.com/larksuite/botframework-go/SDK/message"
	"github.com/larksuite/botframework-go/SDK/protocol"
)

// should set your own appID and appSecret
const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"
)

func main() {
	bot := lark.NewChatBot(appID, appSecret)
	bot.GetTenantAccessTokenInternal(true)

	// example from: https://github.com/larksuite/botframework-go/blob/master/SDK/message/card_builder_test.go
	builder := &message.CardBuilder{}
	// add header
	content := "Please choose color"
	line := 1
	title := protocol.TextForm{
		Tag:     protocol.PLAIN_TEXT_E,
		Content: &content,
		Lines:   &line,
	}
	builder.AddHeader(title, "")
	builder.AddHRBlock()
	// add config
	config := protocol.ConfigForm{
		MinVersion:     protocol.VersionForm{},
		WideScreenMode: true,
	}
	builder.SetConfig(config)
	// add block
	builder.AddDIVBlock(nil, []protocol.FieldForm{
		*message.NewField(false, message.NewMDText("**Async**", nil, nil, nil)),
	}, nil)
	payload1 := make(map[string]string, 0)
	payload1["color"] = "red"
	payload2 := make(map[string]string, 0)
	payload2["color"] = "black"
	builder.AddActionBlock([]protocol.ActionElement{
		message.NewButton(message.NewMDText("red", nil, nil, nil),
			nil, nil, payload1, protocol.PRIMARY, nil, "asyncButton"),
		message.NewButton(message.NewMDText("black", nil, nil, nil),
			nil, nil, payload2, protocol.PRIMARY, nil, "asyncButton"),
	})
	card, _ := builder.Build()

	msg := lark.NewMsgBuffer(lark.MsgInteractive)
	om := msg.BindEmail("zhangwanlong@bytedance.com").CardV4(string(card)).Build()
	bot.PostMessage(om)
}
