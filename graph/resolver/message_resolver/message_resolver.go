package message_resolver

import (
	"server.notion.a1007.wuchuheng.com/m/v2/graph/model"
	"sync"
)

type MessageResolver struct {
}

var messageQueue = struct {
	mu                sync.Mutex
	idMapMessageQueue map[string]*chan *model.Message
}{
	mu:                sync.Mutex{},
	idMapMessageQueue: make(map[string]*chan *model.Message),
}
