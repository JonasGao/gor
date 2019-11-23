package gor

import (
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

type Data struct {
	method string
	key    string
	host   string
	port   string
}

func getData() Data {
	return Data{
		method: os.Getenv("RMETHOD"),
		key:    os.Getenv("RKEY"),
		host:   os.Getenv("RHOST"),
		port:   os.Getenv("RPORT"),
	}
}

func TestNewSSRClient(t *testing.T) {
	data := getData()
	cipher, err := NewStreamCipher(data.method, data.key)
	if err != nil {
		panic(err)
	}
	dialer := net.Dialer{
		Timeout:   time.Millisecond * 500,
		DualStack: true,
	}
	conn, err := dialer.Dial("tcp", net.JoinHostPort(data.host, data.port))
	defer func() {
		err = conn.Close()
		if err != nil {
			t.Error("on close", err)
		}
	}()
	ssconn := NewSSTCPConn(conn, cipher)
	rs := strings.Split(ssconn.RemoteAddr().String(), ":")
	t.Log(rs)
}
