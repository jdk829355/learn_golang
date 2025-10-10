package tcp_conn

import (
	"net"
	"syscall"
	"testing"
	"time"
)

func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			// 임의로 타임아웃을 일으키는 코드
			// Control 콜백은 network connection 직후, 실제 dialoging 이전에 실행
			return &net.DNSError{
				Err:         "connection timed out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

func TestDialTimeout(t *testing.T) {
	// net.DialTimeout은 net.Dialer에 대한 별도의 제어권을 제공하지 않기 때문에
	c, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err == nil {
		c.Close()
		t.Fatal("no timeout")
	}

	// 타입 어설션
	nErr, ok := err.(net.Error)
	if !ok {
		t.Fatal(err)
	}
	if !nErr.Timeout() {
		t.Fatal("error is not a timeout!")
	}
}
