package main

import (
	"errors"
	"fmt"
	"github.com/andy-y-li/socket-go"
	"io"
	"math"
	"net"
	"strconv"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8888"
	DELIMITER      = '\t'
)

func convertToInt32(str string) (int32, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		socketio.PrintLog(fmt.Sprintf("Parse Error: %s\n", err))
		return 0, errors.New(fmt.Sprintf("'%s' is not integer!", str))
	}

	if num > math.MaxInt32 || num < math.MinInt32 {
		socketio.PrintLog(fmt.Sprintf("Convert Error: The integer %s is too large/small.\n", num))
		return 0, errors.New(fmt.Sprintf("'%s' is not 32-bit integer!", num))
	}
	return int32(num), nil
}

func cbrt(param int32) float64 {
	return math.Cbrt(float64(param))
}

func serverGo() {
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		socketio.PrintLog("Listen Error: %s\n", err)
		return
	}
	defer listener.Close()
	socketio.PrintLog("Got listener for the server. (local address: %s)\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			socketio.PrintLog("Accept Error: %s\n", err)
		}
		socketio.PrintLog("Established a connection with a client application. (remote address: %s)\n", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := socketio.Read(conn)
		if err != nil {
			if err == io.EOF {
				socketio.PrintLog("The connection is closed by another side. (Server)\n")
			} else {
				socketio.PrintLog("Read Error: %s (Server)\n", err)
			}
			break
		}
		socketio.PrintLog("Received request: %s (Server)\n", strReq)
		i32Req, err := convertToInt32(strReq)
		if err != nil {
			n, err := socketio.Write(conn, err.Error())
			if err != nil {
				socketio.PrintLog("Write Error (written %d bytes): %s (Server)\n", err)
			}
			socketio.PrintLog("Sent response (written %d bytes): %s (Server)\n", n, err)
			continue
		}
		f64Resp := cbrt(i32Req)
		respMsg := fmt.Sprintf("The cube root of %d is %f.", i32Req, f64Resp)
		n, err := socketio.Write(conn, respMsg)
		if err != nil {
			socketio.PrintLog("Write Error: %s (Server)\n", err)
		}
		socketio.PrintLog("Sent response (written %d bytes): %s (Server)\n", n, respMsg)
	}
}

func main() {
	serverGo()
}
