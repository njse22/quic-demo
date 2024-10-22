package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"github.com/quic-go/quic-go"
	"log"
)

const (
	peerAddress = "192.168.88.100:4244"
	message     = "Hello, multiple streams handling server"
	green       = "\033[97;42m"
	reset       = "\033[0m"
)

func main() {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"multiple-streams-quic-demo"},
	}
	conn, err := quic.DialAddr(context.Background(), peerAddress, tlsConfig, nil)
	if err != nil {
		panic(err)
	}
	for {
		var s byte
		fmt.Printf("\n%sEnter Ctrl-C to quit and any else to continue%s", green, reset)
		_, _ = fmt.Scanf("%d", &s)

		stream, err := conn.OpenStream()
		if err != nil {
			log.Fatalln(err)
		}
		_, err = stream.Write([]byte(message))
		if err != nil {
			log.Fatalln(err)
		}

		buf := make([]byte, len(message))
		_, err = stream.Read(buf)
		if err != nil {
			log.Fatalln(err)
		}

		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					fmt.Println("Dirección IP:", ipNet.IP.String())
				}
			}
		}

		fmt.Printf("%sGot '%s' from stream %d.", reset, buf, stream.StreamID())
	}
}

