package timingwheel

import (
	"container/heap"
	"sync"
	"sync/atomic"
	"time"
)

type item struct {
	Value    interface{}
	Priority int64
	Index    int
}

type priorityQueue []*item

func newPriorityQueue(capacity int) priorityQueue {
	return make(priorityQueue, 0, capacity)
}

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index, pq[j].Index = i, j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	c := cap(*pq)
	if n+1 > c {
		npq := make(priorityQueue, n, c*2)
		copy(npq, *pq)
		*pq = npq
	}
	*pq = (*pq)[0 : n+1]
	item := x.(*item)
	item.Index = n
	(*pq)[n] = item
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	c := cap(*pq)
	if n < (c/2) && c > 25 {
		npq := make(priorityQueue, n, c/2)
		copy(npq, *pq)
		*pq = npq
	}

	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func (pq *priorityQueue) PeekAndShift(max int64) (*item, int64) {
	if pq.Len() == 0 {
		return nil, 0
	}

	item := (*pq)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}
	heap.Remove(pq, 0)
	return item, 0
}

type delayQueue struct {
	C chan *bucket

	mu sync.Mutex
	pq priorityQueue

	sleeping int32
	wakeupC  chan struct{}

	rescheduling int32
	readyC       chan struct{}
}

func newDelayQueue(size int) *delayQueue {
	return &delayQueue{
		C:       make(chan *bucket),
		pq:      newPriorityQueue(size),
		wakeupC: make(chan struct{}),
		readyC:  make(chan struct{}),
	}
}

func (dq *delayQueue) Offer(b *bucket) {
	item := &item{Value: b, Priority: b.Expiration()}

	dq.mu.Lock()
	heap.Push(&dq.pq, item)
	dq.mu.Unlock()

	if item.Index == 0 {
		if atomic.CompareAndSwapInt32(&dq.sleeping, 1, 0) {
			dq.wakeupC <- struct{}{}
		}
		if atomic.CompareAndSwapInt32(&dq.rescheduling, 1, 0) {
			dq.readyC <- struct{}{}
		}
	}
}

func (dq *delayQueue) Poll(exitC chan struct{}) {
	for {
		now := timeToMs(time.Now())

		dq.mu.Lock()
		item, delta := dq.pq.PeekAndShift(now)
		dq.mu.Unlock()

		if item == nil {
			if delta == 0 {
				atomic.StoreInt32(&dq.rescheduling, 1)
				select {
				case <-dq.readyC:
					continue
				case <-exitC:
					goto exit
				}
			} else if delta > 0 {
				atomic.StoreInt32(&dq.sleeping, 1)
				select {
				case <-dq.wakeupC:
					continue
				case <-time.After(time.Duration(delta) * time.Millisecond):
					if atomic.SwapInt32(&dq.sleeping, 0) == 0 {
						<-dq.wakeupC
					}
					continue
				case <-exitC:
					goto exit
				}
			}
		}

		b := item.Value.(*bucket)
		select {
		case dq.C <- b:

		case <-exitC:
			goto exit
		}
	}
exit:
	atomic.StoreInt32(&dq.sleeping, 0)
	atomic.StoreInt32(&dq.rescheduling, 0)
}
