package timer

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

func Now() int64 {
	return time.Now().UnixMilli()
}

var idSet = sync.Map{}

func generateId() int64 {
	randId := rand.Int63()
	for {
		_, ok := idSet.LoadOrStore(randId, struct{}{})
		if !ok {
			return randId
		}
		randId = rand.Int63()
	}
}

type Job struct {
	id int64
	action func()
}

func NewJob(job func()) Job {
	id := generateId()

	return Job{
		id: id,
		action: job,
	}
}

func (j *Job) Id() int64 {
	return j.id
}

type JobList struct {
	lock sync.RWMutex
	list []Job
	pool *ants.Pool
}

func (jl *JobList) AddJob(job Job) {
	jl.lock.Lock()
	defer jl.lock.Unlock()
	jl.list = append(jl.list, job)
}

func (jl *JobList) Clear() {
	jl.lock.Lock()
	defer jl.lock.Unlock()
	jl.list = jl.list[:0]
}

func (jl *JobList) Remove(id int64) {
	jl.lock.Lock()
	defer jl.lock.Unlock()
	for i, j := range jl.list {
		if j.id == id {
			jl.list = append(jl.list[:i], jl.list[i+1:]...)
			break
		}
	}
}

func (jl *JobList) Run() {
	jl.lock.RLock()
	defer jl.lock.RUnlock()
	for _, job := range jl.list {
		jl.pool.Submit(job.action)
	}
}

type Timer struct {
	total time.Duration
	unit time.Duration
	cellCount int64
	jobList []*JobList
	pool *ants.Pool
	ticker atomic.Pointer[time.Ticker]
}

func NewTimer(totalDuration time.Duration, cellCount int64, runnerPoolSize int) (timer *Timer, panicChannel <-chan any, err error) {
	panicChan := make(chan any, 128)
	p, err := ants.NewPool(runnerPoolSize, ants.WithExpiryDuration(totalDuration), ants.WithNonblocking(false), ants.WithPreAlloc(false), ants.WithPanicHandler(func(i interface{}) {
		panicChan <- i
	}))
	if err != nil {
		return nil, nil, err
	}

	unitDurtaion := totalDuration / time.Duration(cellCount)
	jobList := make([]*JobList, cellCount)
	for i := int64(0); i < cellCount; i++ {
		jobList[i] = &JobList{
			list: make([]Job, 0),
			pool: p,
		}
	}


	t := &Timer{
		total: totalDuration,
		unit: unitDurtaion,
		cellCount: cellCount,
		jobList: jobList,
		pool: p,
	}
	return t, panicChan, nil
}

func (t *Timer) AddJob(job Job, cellOrder int64) error {
	if cellOrder < 0 || cellOrder >= int64(len(t.jobList)) {
		return errors.New("cellOrder out of range")
	}
	t.jobList[cellOrder].AddJob(job)
	return nil
}

func (t *Timer) RemoveJob(jobId int64) {
	for _, jl := range t.jobList {
		jl.Remove(jobId)
	}
}

func (t *Timer) Start(ctx context.Context) error {
	if t.ticker.Load() != nil {
		return errors.New("timer already started")
	}

	ticker := time.NewTicker(t.unit)
	t.ticker.Store(ticker)
	go func() {
		cell := int64(0)
		done := ctx.Done()
		for {
			select {
			case <-done:
				t.pool.Release()
				ticker.Stop()
				t.ticker.Store(nil)
				return
			case <-ticker.C:
				t.jobList[cell].Run()
				cell = (cell+1) % t.cellCount
			}
		}
	}()

	return nil
}
