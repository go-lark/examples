package main

import (
	"net/http"
	"testing"

	"github.com/go-lark/lark"
)

func TestFileSize(t *testing.T) {
	if !checkImageSize("./images/fixture.jpg") {
		t.Error("size overflow")
	}
}

func TestMsg(t *testing.T) {
	message := lark.EventMessage{
		Timestamp: "",
		Token:     "",
		EventType: "event_callback",
		Event: lark.EventBody{
			Type:          "message",
			ChatType:      "private",
			MsgType:       "text",
			OpenID:        "ou_8e9e487c3af456fe49aa1b43ffe4ff2e",
			OpenChatID:    "oc_92aeb9f56072dc58cadfd9c1673a7f17",
			Text:          "foosball",
			RealText:      "foosball",
			Title:         "",
			OpenMessageID: "",
			ImageKey:      "",
			ImageURL:      "",
		},
	}
	lark.PostEvent(http.DefaultClient, "http://localhost:8044/lark", message)
}
