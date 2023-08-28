package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	//MR1S()
	//NS1R()
	//MRNS()
	//MR1SS()
	NSS()
}

// M个接收者，1个发送者，发送者通过关闭数据通道表示“不再发送”
// 使用waitgroup来控制数量和关闭channel
func MR1S() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// the sender
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// The only sender can close the
				// channel at any time safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// Receive values until dataCh is
			// closed and the value buffer queue
			// of dataCh becomes empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

/*尽早退出：第一个 select 语句用于尽早退出 goroutine。在每次循环开始时，它尝试从 stopCh 通道中接收数据。
如果 stopCh 通道被关闭（即有其他地方关闭了 stopCh），那么该 select 语句将成功接收数据，从而使得该 goroutine 早早退出。
这样可以确保在外部的 stopCh 通道被关闭后，发送 goroutine 可以尽快退出，避免无谓的发送操作。
防止阻塞：第二个 select 语句用于在发送数据到 dataCh 通道时避免阻塞。
在每次循环开始时，它尝试在两个通道中选择一个可用的操作：接收 stopCh 或发送数据到 dataCh。
这样即使 stopCh 通道被关闭了，如果 dataCh 通道也是可写的，发送操作仍然可以继续，不会因为 stopCh 关闭而导致阻塞。*/

// 一个接收者，N个发送者，唯一的接收者通过关闭额外的信号通道说“请停止发送更多”
// 这是一种比上面的情况稍微复杂一点的情况。我们不能让接收方关闭数据通道来停止数据传输，因为这样做会破坏通道关闭原则。但我们可以让接收方关闭一个额外的信号通道来通知发送方停止发送值。
func NS1R() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel
	// dataCh, and its receivers are the
	// senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// The try-receive operation is to try
				// to exit the goroutine as early as
				// possible. For this specified example,
				// it is not essential.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in the second select may be
				// still not selected for some loops if
				// the send to dataCh is also unblocked.
				// But this is acceptable for this
				// example, so the first select block
				// above can be omitted.
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				// The receiver of channel dataCh is
				// also the sender of stopCh. It is
				// safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}

// M 个接收者，N 个发送者，其中任何一个通过通知主持人关闭附加信号通道来表示“让我们结束游戏”
func MRNS() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown
	// below, and its receivers are all senders
	// and receivers of dataCh.
	toStop := make(chan string, 1)
	// The channel toStop is used to notify the
	// moderator to close the additional signal
	// channel (stopCh). Its senders are any senders
	// and receivers of dataCh, and its receiver is
	// the moderator goroutine shown below.
	// It must be a buffered channel.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					// Here, the try-send operation is
					// to notify the moderator to close
					// the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// The try-receive operation here is to
				// try to exit the sender goroutine as
				// early as possible. Try-receive and
				// try-send select blocks are specially
				// optimized by the standard Go
				// compiler, so they are very efficient.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and for ever in theory) if the send
				// to dataCh is also non-blocking. If
				// this is unacceptable, then the above
				// try-receive operation is essential.
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// Same as the sender goroutine, the
				// try-receive operation here is to
				// try to exit the receiver goroutine
				// as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				// Even if stopCh is closed, the first
				// branch in this select block might be
				// still not selected for some loops
				// (and forever in theory) if the receive
				// from dataCh is also non-blocking. If
				// this is not acceptable, then the above
				// try-receive operation is essential.
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						// Here, the same trick is
						// used to notify the moderator
						// to close the additional
						// signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

// M 个接收者，一个发送者”情况的一种变体：关闭请求由第三方 goroutine 发出
func MR1SS() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)
	closing := make(chan struct{}) // signal channel
	closed := make(chan struct{})

	// The stop function can be called
	// multiple times safely.
	stop := func() {
		select {
		case closing <- struct{}{}:
			<-closed
		case <-closed:
		}
	}

	// some third-party goroutines
	for i := 0; i < NumThirdParties; i++ {
		go func() {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop()
		}()
	}

	// the sender
	go func() {
		defer func() {
			close(closed)
			close(dataCh)
		}()

		for {
			select {
			case <-closing:
				return
			default:
			}

			select {
			case <-closing:
				return
			case dataCh <- rand.Intn(Max):
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

// “N个发送方”情况的一种变体：必须关闭数据通道才能告诉接收方数据发送结束
func NSS() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 1000000
	const NumReceivers = 10
	const NumSenders = 1000
	const NumThirdParties = 15

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)     // will be closed
	middleCh := make(chan int)   // will never be closed
	closing := make(chan string) // signal channel
	closed := make(chan struct{})

	var stoppedBy string

	// The stop function can be called
	// multiple times safely.
	stop := func(by string) {
		select {
		case closing <- by:
			<-closed
		case <-closed:
		}
	}

	// the middle layer
	go func() {
		exit := func(v int, needSend bool) {
			close(closed)
			if needSend {
				dataCh <- v
			}
			close(dataCh)
		}

		for {
			select {
			case stoppedBy = <-closing:
				exit(0, false)
				return
			case v := <-middleCh:
				select {
				case stoppedBy = <-closing:
					exit(v, true)
					return
				case dataCh <- v:
				}
			}
		}
	}()

	// some third-party goroutines
	for i := 0; i < NumThirdParties; i++ {
		go func(id string) {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop("3rd-party#" + id)
		}(strconv.Itoa(i))
	}

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					stop("sender#" + id)
					return
				}

				select {
				case <-closed:
					return
				default:
				}

				select {
				case <-closed:
					return
				case middleCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for range [NumReceivers]struct{}{} {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
