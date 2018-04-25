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
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func Write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}
