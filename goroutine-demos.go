package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// ---------------------------------------------------------------------------
// 1. Gosched

func testGoSched() {
	for i := 0; i < 5; i++ {
		go func(val int) {
			for i := 0; i < 3; i++ {
				if val == 3 {
					// Gosched yields the processor, allowing other goroutines to run. It does not
					// suspend the current goroutine, so execution resumes automatically.
					// 这个函数的作用是让当前goroutine让出CPU，好让其它的goroutine获得执行的机会。
					// 同时，当前的goroutine也会在未来的某个时间点继续运行。
					runtime.Gosched()
				}
				fmt.Println(val)
			}
		}(i)
	}

	time.Sleep(time.Second * 3)
	fmt.Println("End")
}

// ---------------------------------------------------------------------------
// 2. GOMAXPROCS

func testGOMAXPROCS() {
	// 设置可同时执行的逻辑Cpu数量，默认和硬件的线程数一致而不是核心数，可以通过调用GOMAXPROCS(-1)来获取当前逻辑Cpu数。
	// <1：不修改任何数值。
	// =1：单核心执行。
	// >1：多核并发执行。
	fmt.Println(runtime.NumCPU())       // 4
	fmt.Println(runtime.GOMAXPROCS(-1)) // 4
	fmt.Println(runtime.GOMAXPROCS(1))  // 4
	fmt.Println(runtime.GOMAXPROCS(2))  // 1
	fmt.Println(runtime.GOMAXPROCS(8))  // 2
	fmt.Println(runtime.GOMAXPROCS(-1)) // 8
}

// ---------------------------------------------------------------------------
// 3. Goexit

func testGoexit() {
	// Output:
	//  End2
	//  End1
	go func() {
		defer func() {
			fmt.Println("End2")
		}()

		// 1. Goexit立即终止当前协程，不会影响其它协程。终止前会调用此协程声明的defer方法。
		// 由于Goexit不是panic，所以recover捕获的error会为nil。
		// 2.	当main方法所在主协程调用Goexit时，Goexit不会return，所以主协程将继续等待子协程执行，
		// 当所有子协程执行完时，程序报错deadlock
		runtime.Goexit()
		fmt.Println("End3")
	}()

	time.Sleep(time.Second)
	fmt.Println("End1")
}

// ---------------------------------------------------------------------------
// 4. 无缓存channel

func testChannel() {
	// 创建无缓冲管道
	ch := make(chan int) // ch : make(chan int, 0)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i + 1)
			time.Sleep(time.Second)
		}
		ch <- 3
	}()

	// 1. 该线程阻塞到这里，当没有其他线程写数据的时候，这个读操作会阻塞到这里。
	x := <-ch
	fmt.Println("Rev: ", x)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i + 1)
			time.Sleep(time.Second)
		}
		n := <-ch
		fmt.Println(n)
	}()

	// 2. 该线程会阻塞到这里，当没有其他线程读数据的时候，这个写操作会阻塞到这里。
	ch <- 10
	fmt.Println("Write")
}

// ---------------------------------------------------------------------------
// 5. 有缓存channel

func testBufferedChannel() {
	// 带缓存的channel类似一个线程安全的队列。
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i + 1)
			time.Sleep(time.Second)
		}
		n := <-ch
		fmt.Println("Read:", n)

		n = <-ch
		fmt.Println("Read:", n)

		n = <-ch
		fmt.Println("Read:", n)

		// 2. 读操作阻塞到这里，因为缓存区没有数据，直到有数据被写。
		n = <-ch
		fmt.Println("Read:", n)
	}()

	ch <- 100
	ch <- 101
	// 1. 缓存区已满，写操作会阻塞到这里，直到缓存区有数据被取走。
	ch <- 102

	for i := 0; i < 3; i++ {
		fmt.Println(i + 1)
		time.Sleep(time.Second)
	}

	ch <- 103
	time.Sleep(time.Second)
	fmt.Println("end")
}

// ---------------------------------------------------------------------------
// 6. select

func testSelect() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2

	ch2 := make(chan int, 2)
	ch2 <- 3
	ch2 <- 4

	for {
		exit := false

		// select会随机选择一个没有阻塞的分支执行，
		// 当所有分支都阻塞后则进入default分支。
		select {
		case x := <-ch1:
			fmt.Println(x)
		case x := <-ch2:
			fmt.Println(x)
		default:
			exit = true
		}

		if exit {
			break
		}
	}
}

// ---------------------------------------------------------------------------
// 6. timer

func testTimer() {
	// 创建并启动一个timer
	timer1 := time.NewTimer(time.Second * 5)
	fmt.Println(time.Now())
	end := <-timer1.C
	fmt.Println(end)

	timer1.Reset(time.Second * 100)
	timer1.Stop()

	// time.After封装了time.NewTimer
	end = <-time.After(time.Second * 5)
	fmt.Println(end)

	// 周期性定时器
	ticker := time.NewTicker(time.Second * 2)
	for i := 0; i < 5; i++ {
		n := <-ticker.C
		fmt.Println(n)

		if i == 2 {
			ticker.Stop()
			break
		}
	}
}

// ---------------------------------------------------------------------------
// 6. sync.WaitGroup

func testWaitGroup() {
	// WaitGroup类似CountDownLatch这样的同步计数器。
	// WaitGroup 对象内部有一个计数器，最初从0开始。Add(n)增加计数值，
	// Done()每次把计数器-1 ，wait()会阻塞代码的运行，直到计数器地值减为0。
	// 注意：WaitGroup对象不是一个引用类型，当函数参数要传址。
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go func(n int) {
			time.Sleep(time.Second * 2)
			fmt.Println(n)
			wg.Done()
		}(i)

		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("End")
}

// ---------------------------------------------------------------------------
// 7.sync.Mutex

func testMutex() {
	num := 100

	mut := sync.Mutex{}

	go func() {
		for {
			mut.Lock()

			if num <= 0 {
				break
			}
			num = num - 5
			time.Sleep(time.Microsecond * 10)
			fmt.Println("Go1", num)

			mut.Unlock()
		}
	}()

	go func() {
		for {
			mut.Lock()

			if num <= 0 {
				break
			}
			num = num - 5
			time.Sleep(time.Microsecond * 10)
			fmt.Println("Go2", num)

			mut.Unlock()
		}
	}()

	time.Sleep(time.Second * 2)
}

// ---------------------------------------------------------------------------
// 8. sync.RWMutex

// 可以多个goroutine同时读。
// 写的时候，其他goroutine不能读也不能写。
func testRWMutex() {
	wg := sync.WaitGroup{}
	rwMutex := sync.RWMutex{}

	t1 := time.Now()
	for i := 0; i < 5; i++ {
		go func() {
			rwMutex.RLock()
			time.Sleep(time.Second * 2)
			rwMutex.RUnlock()
			wg.Done()
		}()

		wg.Add(1)
	}
	wg.Wait()

	duration := time.Now().Sub(t1)
	// 输出：2.0007317。说明多个goroutine同时进行。
	fmt.Println(duration.Seconds())

	t1 = time.Now()

	for i := 0; i < 5; i++ {
		go func() {
			rwMutex.Lock()
			time.Sleep(time.Second * 2)
			rwMutex.Unlock()
			wg.Done()
		}()

		wg.Add(1)
	}
	wg.Wait()

	duration = time.Now().Sub(t1)
	// 输出：10.0018662。说明同时只有一个goroutine进行。
	fmt.Println(duration.Seconds())
}

// ---------------------------------------------------------------------------
// 9. sync.Once

// 类似c++中的std::call_once。可以保证某个函数在并发中只执行一次。
func testCallOnce() {
	once := sync.Once{}

	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("Call once")
			})
		}()
	}
}

// ---------------------------------------------------------------------------
// 10. sync.Cond

// 条件变量和c++中的std::condition_variable使用类似，需要配合锁一起使用。
func testConditionVariable() {
	cv := sync.NewCond(&sync.Mutex{})
	flag := true

	go func() {
		time.Sleep(time.Second * 3)
		cv.L.Lock()
		time.Sleep(time.Second * 3)
		flag = false
		cv.Signal()
		cv.L.Unlock()
	}()

	cv.L.Lock()
	for flag { // 防止假唤醒
		fmt.Println("Wait")
		cv.Wait()
	}
	fmt.Println("End")
	cv.L.Unlock()
}

// ---------------------------------------------------------------------------
// 11. sync/atomic

func testAtomic() {
	var val int32 = 1

	// 原子读取一个变量的值，否则读取一个变量的时候可能被其他线程打断。
	n := atomic.LoadInt32(&val)
	fmt.Println(n)

	atomic.AddInt32(&val, 2)
	fmt.Println(val)

	// 原子地对一个变量赋值。
	atomic.StoreInt32(&val, 3)
	fmt.Println(val)

	// 原子地对一个变量赋值，并且返回旧值。
	atomic.SwapInt32(&val, 4)
	fmt.Println(val)

	// 将val和old值对比，相等时则和新值进行交换。
	s := atomic.CompareAndSwapInt32(&val, 4, 6)
	fmt.Println(s, val)
}

// ---------------------------------------------------------------------------
// 12. deadlock

func testDeadLock() {
	// 可以输出 1
	// fatal error: all goroutines are asleep - deadlock
	fmt.Println("1")
	ch := make(chan int)
	ch <- 0
	<-ch // 程序根本执行不到这条语句，造成死锁crash
}

// 经典死锁案例，哲学家就餐。
func testDeadLock2() {
	// fatal error: all goroutines are asleep - deadlock!
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 次goroutine等待ch1，当从ch1中读到数据的时候才会向ch2中写数据。
	// 主goroutine等待ch2，当从ch2中读到数据的时候才会向ch1中写数据。

	go func() {
		<-ch1
		ch2 <- 1
	}()

	<-ch2
	ch1 <- 1

	time.Sleep(time.Second)
}

// ---------------------------------------------------------------------------

func main() {
}
