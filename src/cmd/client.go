package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/src/znet"
)

func main() {
	fmt.Println("[Client] client start")

	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("[Client] client dial error:", err)
		return
	}

	dp := znet.NewDataPack()
	fmt.Println("[Client] client dial ok")

	msg := znet.NewMassage(1, []byte("Trigger panic test"))
	data, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("[Client] pack error:", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("[Client] write error:", err)
		return
	}

	fmt.Println("[Client] sent trigger message")

	time.Sleep(2 * time.Second)

	headData := make([]byte, dp.GetHeadLen())
	n, err := io.ReadFull(conn, headData)
	if err != nil {
		fmt.Printf("[Client] read head error: %v, n=%d\n", err, n)
	} else {
		msgHead, err := dp.UnPack(headData)
		if err == nil {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			_, err = io.ReadFull(conn, msg.Data)
			if err == nil {
				fmt.Println("[Client] received:", string(msg.GetData()))
			}
		}
	}

	conn.Close()
	fmt.Println("[Client] client exit")
}
