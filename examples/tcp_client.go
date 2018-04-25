package main

import (
	"fmt"
	"github.com/andy-y-li/socket-go"
	"io"
	"math/rand"
	"net"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8888"
)

func clientGo(id int) {

	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		socketio.PrintLog("Dial Error: %s (Client[%d])\n", err, id)
		return
	}
	defer conn.Close()
	socketio.PrintLog("Connected to server. (remote address: %s, local address: %s) (Client[%d])\n", conn.RemoteAddr(), conn.LocalAddr(), id)

	time.Sleep(200 * time.Millisecond)

	requestTimes := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	for i := 0; i < requestTimes; i++ {
		i32Req := rand.Int31()
		n, err := socketio.Write(conn, fmt.Sprintf("%d", i32Req))
		if err != nil {
			socketio.PrintLog("Write Error: %s (Client[%d])\n", err, id)
			continue
		}
		socketio.PrintLog("Sent request (written %d bytes): %d (Client[%d])\n", n, i32Req, id)
	}
	for j := 0; j < requestTimes; j++ {
		strResp, err := socketio.Read(conn)
		if err != nil {
			if err == io.EOF {
				socketio.PrintLog("The connection is closed by another side. (Client[%d])\n", id)
			} else {
				socketio.PrintLog("Read Error: %s (Client[%d])\n", err, id)
			}
			break
		}
		socketio.PrintLog("Received response: %s (Client[%d])\n", strResp, id)
	}
}

func main() {
	clientGo(1)
}
