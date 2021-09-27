// package golang_demo
// name concurrency

package golang_demo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func serveApp(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		// wait for the stop signal
		<-stop
		// shutdown 即使正在监听可以使应用平滑退出 ListenAndServe 对标的操作 Shutdown
		s.Shutdown(context.Background())
		fmt.Println("serveApp shutdown")
	}()
	return s.ListenAndServe()
}

func serveDebug(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		// wait for the stop signal
		<-stop
		// shutdown 即使正在监听可以使应用平滑退出 ListenAndServe 对标的操作 Shutdown
		s.Shutdown(context.Background())
		fmt.Println("serveDebug shutdown")
	}()
	return s.ListenAndServe()
}

func Concurrency() {
	done := make(chan error, 2) // 两个应用，接收 err
	stop := make(chan struct{}) // 0 值 channel

	timer := time.NewTimer(2 * time.Second)

	go func() {
		done <- serveApp("127.0.0.1:25600", nil, stop)
	}()

	go func() {
		done <- serveDebug("127.0.0.1:25601", nil, stop)
	}()

	isStopped := false

	go func() {
		select {
		case <-timer.C:
			if !isStopped {
				close(stop)
				isStopped = true
			}
		}
	}()

	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("err:", err)
		}

		if !isStopped {
			// 唤醒后台两个 goroutine
			close(stop)
			isStopped = true
		}
	}
}
