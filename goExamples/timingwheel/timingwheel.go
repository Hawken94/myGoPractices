package timingwheel

import (
	"errors"
	"sync/atomic"
	"time"
	"unsafe"
)

type TimingWheel struct {
	tick      int64
	wheelSize int64

	interval    int64
	currentTime int64
	buckets     []*bucket
	queue       *delayQueue

	overflowWheel unsafe.Pointer

	workerPoolSize int64

	exitC     chan struct{}
	waitGroup waitGroupWrapper
}

func NewTimingWheel(tick time.Duration, wheelSize int64) *TimingWheel {
	tickMs := int64(tick / time.Millisecond)
	if tickMs <= 0 {
		panic(errors.New("tick must be greater than or equal to 1ms"))
	}
	startMs := timeToMs(time.Now())
	return newTimingWheel(
		tickMs,
		wheelSize,
		startMs,
		newDelayQueue(int(wheelSize)),
	)
}
func newTimingWheel(tickMs, wheelSize, startMs int64, queue *delayQueue) *TimingWheel {
	buckets := make([]*bucket, wheelSize)

	for i := range buckets {
		buckets[i] = newBucket()
	}
	return &TimingWheel{
		tick:        tickMs,
		wheelSize:   wheelSize,
		currentTime: truncate(startMs, tickMs),
		interval:    tickMs * wheelSize,
		buckets:     buckets,
		queue:       queue,
		exitC:       make(chan struct{}),
	}
}

func (tw *TimingWheel) add(t *Timer) bool {
	if t.expiration < tw.currentTime+tw.tick {
		return false
	} else if t.expiration < tw.currentTime+tw.interval {
		virtualID := t.expiration / tw.tick
		b := tw.buckets[virtualID%tw.wheelSize]
		b.Add(t)

		if b.setExpiration(virtualID * tw.tick) {
			tw.queue.Offer(b)
		}

		return true
	} else {
		overflowWheel := atomic.LoadPointer(&tw.overflowWheel)
		if overflowWheel == nil {
			atomic.CompareAndSwapPointer(
				&tw.overflowWheel,
				nil,
				unsafe.Pointer(newTimingWheel(
					tw.interval,
					tw.wheelSize,
					tw.currentTime,
					tw.queue,
				)),
			)
			overflowWheel = atomic.LoadPointer(&tw.overflowWheel)
		}
		return (*TimingWheel)(overflowWheel).add(t)
	}

}

func (tw *TimingWheel) addOrRun(t *Timer) {
	if !tw.add(t) {
		go t.task()
	}
}

func (tw *TimingWheel) advanceClock(expiration int64) {
	if expiration >= tw.currentTime+tw.tick {
		tw.currentTime = truncate(expiration, tw.tick)

		overflowWheel := atomic.LoadPointer(&tw.overflowWheel)
		if overflowWheel != nil {
			(*TimingWheel)(overflowWheel).advanceClock(tw.currentTime)
		}
	}
}

func (tw *TimingWheel) Start() {
	tw.waitGroup.Wrap(func() {
		tw.queue.Poll(tw.exitC)
	})

	tw.waitGroup.Wrap(func() {
		for {
			select {
			case b := <-tw.queue.C:
				tw.advanceClock(b.Expiration())
				b.Flush(tw.addOrRun)
			case <-tw.exitC:
				return
			}
		}
	})
}

func (tw *TimingWheel) Stop() {
	close(tw.exitC)
	tw.waitGroup.Wait()
}

func (tw *TimingWheel) AfterFunc(d time.Duration, f func()) *Timer {
	t := &Timer{
		expiration: timeToMs(time.Now().Add(d)),
		task:       f,
	}
	tw.addOrRun(t)
	return t
}
