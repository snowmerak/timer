package timer

import (
	"context"
	"github.com/Workiva/go-datastructures/queue"
	"sync"
	"time"
)

var milliSecond = time.Millisecond.Milliseconds()

func Now() int64 {
	return time.Now().UnixMilli()
}

type node struct {
	at       int64
	interval time.Duration
	action   func()
}

func (n *node) Compare(other queue.Item) int {
	return int(n.at - other.(*node).at)
}

var nodePool = &sync.Pool{
	New: func() interface{} {
		return &node{}
	},
}

type Timer struct {
	queue         *queue.PriorityQueue
	context       context.Context
	contextCancel context.CancelFunc
	name          string

	lcm     int64
	started bool
}

func NewTimer(ctx context.Context, name string, number int) *Timer {
	t := &Timer{
		context: ctx,
		queue:   queue.NewPriorityQueue(number+1, true),
		name:    name,
	}
	t.context, t.contextCancel = context.WithCancel(t.context)
	t.queue.Put(&node{
		at:       Now(),
		interval: time.Hour,
		action:   func() {},
	})
	return t
}

func (t *Timer) Add(interval time.Duration, action func()) bool {
	if !t.started {
		if t.lcm == 0 {
			t.lcm = interval.Milliseconds()
		} else {
			t.lcm = BinaryGcd(t.lcm, interval.Milliseconds())
		}
	}

	n := nodePool.Get().(*node)
	n.interval = interval
	n.at = Now() + n.interval.Milliseconds()
	n.action = action

	err := t.queue.Put(n)
	if err != nil {
		return false
	}
	return true
}

func (t *Timer) Start() {
	if t.started {
		return
	}
	t.started = true
	go func() {
		now := time.Now().UnixMilli()
		var cache []*node
		for {
			select {
			case <-t.context.Done():
				return
			default:
				now += t.lcm
				for {
					if t.queue.Empty() {
						break
					}
					i, err := t.queue.Get(1)
					if err != nil {
						break
					}
					if i == nil {
						break
					}
					n := i[0].(*node)
					if n.at > now {
						cache = append(cache, n)
						break
					}
					n.action()
					cache = append(cache, n)
				}
				go func() {
					for _, n := range cache {
						t.Add(n.interval, n.action)
						n.action = nil
						nodePool.Put(n)
					}
					cache = nil
				}()
				time.Sleep(time.Duration(t.lcm) * time.Millisecond)
			}
		}
	}()
}

func (t *Timer) Stop() {
	t.contextCancel()
}
