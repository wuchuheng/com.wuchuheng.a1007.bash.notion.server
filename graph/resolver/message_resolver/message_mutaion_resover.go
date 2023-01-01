package message_resolver

import (
	"context"
	"fmt"
	"server.notion.a1007.wuchuheng.com/m/v2/graph/model"
)

var id = 0

func (r *MessageResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	id++
	newMessage := model.Message{
		ID:      fmt.Sprint(id),
		Title:   input.Title,
		Content: input.Content,
	}
	go func() {
		messageQueue.mu.Lock()
		for uid, el := range messageQueue.idMapMessageQueue {
			fmt.Printf("%s\n", uid)
			*el <- &newMessage
		}
		messageQueue.mu.Unlock()
	}()

	return &newMessage, nil

}
