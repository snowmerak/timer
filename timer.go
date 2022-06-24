package timer

import (
	"context"
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"log"
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
	queue   *queue.PriorityQueue
	context context.Context

	lcm     int64
	started bool
}

func NewTimer(ctx context.Context, number int) *Timer {
	return &Timer{
		context: ctx,
		queue:   queue.NewPriorityQueue(number, true),
	}
}

func (t *Timer) Add(interval time.Duration, action func()) bool {
	if !t.started {
		if t.lcm == 0 {
			t.lcm = interval.Milliseconds()
		} else {
			t.lcm = BinaryGcd(t.lcm, interval.Milliseconds())
		}
		fmt.Println(t.lcm)
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
						if err := t.queue.Put(n); err != nil {
							log.Println(err)
						}
						break
					}
					go n.action()
					if ok := t.Add(n.interval, n.action); !ok {
						log.Println("Cannot add previous action")
					}
					n.action = nil
					nodePool.Put(n)
				}
				time.Sleep(time.Duration(t.lcm) * time.Millisecond)
			}
		}
	}()
}

func (t *Timer) Stop() {
	t.context.Done()
}