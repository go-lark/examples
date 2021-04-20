package main

import (
	"github.com/go-lark/lark"
)

// should set your own appID and appSecret
const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"
)

// we recommend use https://open.feishu.cn/tool/cardbuilder to build card body
// and then generate to struct with [Paste JSON as Code](https://app.quicktype.io/)
func main() {
	bot := lark.NewChatBot(appID, appSecret)
	bot.GetTenantAccessTokenInternal(true)
	cardContent := `{
		"config": {
			"wide_screen_mode": false
		},
		"elements": [
			{
				"tag": "div",
				"text": {
					"i18n": {
						"zh_cn": "中文文本",
						"en_us": "English text",
						"ja_jp": "日本語文案"
					},
					"tag": "plain_text"
				}
			},
			{
				"tag": "div",
				"text": {
					"tag": "plain_text",
					"content": "This is a very very very very very very very long text;"
				}
			},
			{
				"actions": [
					{
						"tag": "button",
						"text": {
							"content": "a",
							"tag": "plain_text"
						},
						"type": "default"
					}
				],
				"tag": "action"
			}
		],
		"header": {
			"title": {
				"content": "a",
				"tag": "plain_text"
			}
		}
	}
	`

	msg := lark.NewMsgBuffer(lark.MsgInteractive)
	om := msg.BindEmail("zhangwanlong@bytedance.com").Card(cardContent).Build()
	bot.PostMessage(om)
}
