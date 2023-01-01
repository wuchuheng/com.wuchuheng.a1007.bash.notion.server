package message_resolver

import (
	"context"
	"github.com/google/uuid"
	"server.notion.a1007.wuchuheng.com/m/v2/graph/model"
)

func (r *MessageResolver) NewMessage(ctx context.Context) (<-chan *model.Message, error) {
	newMessageQueue := make(chan *model.Message)
	go func() {
		uid := uuid.New().String()
		messageQueue.mu.Lock()
		messageQueue.idMapMessageQueue[uid] = &newMessageQueue
		messageQueue.mu.Unlock()
		select {
		case <-ctx.Done():
			close(newMessageQueue)
			messageQueue.mu.Lock()
			delete(messageQueue.idMapMessageQueue, uid)
			messageQueue.mu.Unlock()

			return
		}
	}()

	return newMessageQueue, nil
}
