package limit

import (
	"fmt"
	"sync"
)

// Limit 并发阻塞, 安全退出
type Limit struct {
	wg sync.WaitGroup
	ch chan struct{}
}

// New create LimitWaitGroup
func New(opts ...Option) *Limit {
	options := newOptions(opts...)
	return &Limit{
		ch: make(chan struct{}, options.Max),
	}
}

// Add add
func (l *Limit) Add(delta int) {
	l.ch <- struct{}{}
	l.wg.Add(delta)
}

// Done done
func (l *Limit) Done() {
	<-l.ch
	l.wg.Done()
}

//Do execute
func (l *Limit) Do(Func func()) {
	l.Add(1)
	go func() {
		defer l.Done()
		Func()
	}()
}

// Wait wait
func (l *Limit) Wait() {
	l.wg.Wait()
	close(l.ch)
}

// Close close
func (l *Limit) Close() error {
	close(l.ch)
	return nil
}

// Stat limit stat
func (l *Limit) Stat() string {
	return fmt.Sprintf("Limit => max: %d, current: %d, idle: %d", cap(l.ch), len(l.ch), cap(l.ch)-len(l.ch))
}

// Len len
func (l *Limit) Len() int {
	return len(l.ch)
}
