package socketio

import (
	"bytes"
	"fmt"
	"net"
)

const (
	DELIMITER = '\t'
)

var logSn = 1

func PrintLog(format string, args ...interface{}) {
	fmt.Printf("%d: %s", logSn, fmt.Sprintf(format, args...))
	logSn++
}

func Read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 65535)
	//var buffer bytes.Buffer
	//for {
	n, err := conn.Read(readBytes)
	if err != nil {
		return "", err
	}
	return string(readBytes[:n]), nil
	/*
			readByte := readBytes[0]
			if readByte == DELIMITER ||
				readByte == '\r' ||
				readByte == '\n' {
				break
			}
			buffer.WriteByte(readByte)
		//}
	*/
	//return buffer.String(), nil
}

func Write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	//buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}
