package tcp_conn

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	// tcp 바인딩
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	// 클라이언트와 서버를 동기화 할 채널
	done := make(chan struct{})

	go func() {
		// 리스너 고루틴이 반환될 때 done에 송신
		// 이는 main 고루틴을 반환하게 함
		defer func() { done <- struct{}{} }()

		for {
			// 연결 요청 받기
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}
			// 핸들러
			go func(c net.Conn) {
				defer func() {
					c.Close()
					// 리스너 고루틴을 반환시키는 역할
					done <- struct{}{}
				}()
				// 데이터를 수신할 버퍼
				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						// EOF 즉, 읽을 데이터가 더는 없으면 핸들러 고루틴 반환 (FIN 패킷 받았을 때)
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	// 우아한 종료 시작
	conn.Close()

	// 연결을 처리하는 핸들러 고루틴이 반환되며 done에 데이터를 송신하여 listener를 close
	<-done
	listener.Close()
	// 리스너가 close되면서 최종 고루틴을 종료
	<-done
}
