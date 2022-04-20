package protocol

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"net"
	"os"
	"time"
)

type Server interface {
	Listen(port string)
	Close()
}

type serverImpl struct {
	listener net.Listener
}

func (s serverImpl) Listen(port string) {
	fmt.Println("starting server...")

	var err error
	s.listener, err = net.Listen(Protocol, port)
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitConnectionErrorCode)
	}

	fmt.Println("awaiting connections...")
	s.acceptConn()
}

func (s serverImpl) Close() {
	s.listener.Close()
}

func (s serverImpl) acceptConn() {
	var conn net.Conn
	var err error
	for {
		conn, err = s.listener.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(ExitConnectionErrorCode)
		}
	}

	s.handShake(conn)
}

func (s serverImpl) handShake(conn net.Conn) error {
	signatureServerSection := uuid.New().String()
	_, err := conn.Write([]byte(signatureServerSection))
	if err != nil {
		fmt.Printf("could not handshake: %v", err)
		return err
	}
	for start := time.Now(); time.Since(start) < time.Second; {
		signatureClientSection, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("could not handshake: %v", err)
			return err
		}

		_, err = uuid.Parse(signatureClientSection)
		if err != nil {
			fmt.Printf("could not handshake: %v", err)
			return err
		}
	}
}

func NewServer() Server {
	return &serverImpl{}
}
