package utility

import (
	"sync"
	"sync/atomic"
)

type SyncGroup struct {
	sync.WaitGroup
	count int64
}

func NewSyncGroup() *SyncGroup {
	sg := SyncGroup{count: 0}
	return &sg
}

func (sg *SyncGroup) Add(d int) {
	atomic.AddInt64(&sg.count, int64(d))
	sg.WaitGroup.Add(d)
}

func (sg *SyncGroup) Done() {
	atomic.AddInt64(&sg.count, -1)
	sg.WaitGroup.Done()
}

func (sg *SyncGroup) Count() int {
	return int(atomic.LoadInt64(&sg.count))
}
