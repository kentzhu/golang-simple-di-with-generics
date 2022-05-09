package impl

import "simpledi/example/service"

func NewMessageService(messageContent string) service.IMessageService {
	return &messageService{content: messageContent}
}

type messageService struct {
	content string
}

func (m *messageService) Message() string {
	return m.content
}
