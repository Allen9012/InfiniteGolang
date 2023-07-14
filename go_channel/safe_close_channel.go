package main

import "sync"

/*粗暴关闭，解决了问题，但是违反了通道关闭原则*/

func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch)   // panic if ch is closed
	return true // <=> justClosed = true; return
}

func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}

/*优雅关闭*/

// 使用sync.Once来保证只关闭一次

type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

// 使用Mutex来保证只关闭一次
type MyChannel2 struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel2() *MyChannel2 {
	return &MyChannel2{C: make(chan T)}
}
func (mc *MyChannel2) SafeClose() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
}
func (mc *MyChannel2) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
