package main

import (
	"fmt"
	"sync"
	"time"
	"zinx/src/ziface"

	"zinx/src/znet"
)

type Router1 struct {
	znet.BaseRouter
}

func (r *Router1) Handle(req ziface.IRequest) {
	req.GetConnection().SendMsg(1, []byte("hello zinx"))
}

type PanicTestRouter struct {
	znet.BaseRouter
}

func (r *PanicTestRouter) Handle(req ziface.IRequest) {
	conn := req.GetConnection()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			msg := fmt.Sprintf("Concurrent message %d at %v", index, time.Now().UnixNano())
			err := conn.SendMsg(1, []byte(msg))
			if err != nil {
				fmt.Printf("Send error: %v\n", err)
			} else {
				fmt.Printf("Sent: %s\n", msg)
			}
		}(i)

		if i == 50 {
			time.Sleep(5 * time.Millisecond)
			fmt.Println(">>> Stopping connection while sending...")
			conn.Stop()
		}
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}

func main() {
	fmt.Println("[Server] server start")
	s := znet.NewServer()
	s.AddRouter(1, &PanicTestRouter{znet.BaseRouter{}})
	s.Serve()
}

// ... existing code ...
