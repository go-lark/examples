package main_test

import (
	"net/http"
	"testing"

	"github.com/go-lark/lark"
)

func TestSendByID(t *testing.T) {
	message := lark.EventMessage{
		Timestamp: "",
		Token:     "",
		EventType: "event_callback",
		Event: lark.EventBody{
			Type:          "message",
			ChatType:      "private",
			MsgType:       "text",
			OpenID:        "ou_08198ccd6a37644b49f4789c92369c80",
			OpenChatID:    "oc_1c09434c2264eb52dc7667895a7fac6d",
			Text:          "/i 2A5wXdxrWghhSvvFFb",
			Title:         "",
			OpenMessageID: "",
			ImageKey:      "",
			ImageURL:      "",
		},
	}

	lark.PostEvent(http.DefaultClient, "http://localhost:8044/lark", message)
}

func TestSendByAt(t *testing.T) {
	message := lark.EventMessage{
		Timestamp: "",
		Token:     "",
		EventType: "event_callback",
		Event: lark.EventBody{
			Type:          "message",
			ChatType:      "private",
			MsgType:       "text",
			OpenID:        "ou_08198ccd6a37644b49f4789c92369c80",
			OpenChatID:    "oc_1c09434c2264eb52dc7667895a7fac6d",
			Text:          "/s hello",
			Title:         "",
			OpenMessageID: "",
			ImageKey:      "",
			ImageURL:      "",
		},
	}

	lark.PostEvent(http.DefaultClient, "http://localhost:8044/lark", message)
}
