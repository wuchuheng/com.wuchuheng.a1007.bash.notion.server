package message_resolver

import (
	"context"
	"server.notion.a1007.wuchuheng.com/m/v2/graph/model"
)

func (r *MessageResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return messageList, nil
}
