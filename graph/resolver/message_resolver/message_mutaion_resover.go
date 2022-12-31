package message_resolver

import (
	"context"
	"fmt"
	"server.notion.a1007.wuchuheng.com/m/v2/graph/model"
)

func (r *MessageResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	newMessage := model.Message{
		ID:      fmt.Sprint(len(messageList) + 1),
		Title:   input.Title,
		Content: input.Content,
	}
	messageList = append(messageList, &newMessage)

	return &newMessage, nil

}
