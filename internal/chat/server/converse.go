package server

import (
	"github.com/Madslick/chit-chat-go/pkg"
)

func (cs *ChatroomServer) Converse(stream pkg.Chatroom_ConverseServer) error {
	return cs.conversationAssistant.Converse(stream)
}
