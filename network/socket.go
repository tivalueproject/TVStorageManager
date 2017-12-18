package network

import (
	"net"
	"fmt"
)

var bufsize int = 1024*100

func Listen(addr string)(net.Listener, error) {

  // Listen for incoming connections.
  listener, err := net.Listen("tcp", addr)
  if err != nil {
      fmt.Println("Error listening:", err.Error())
      panic(err)
  }

  return listener, err
}

func Read(conn net.Conn) (string, error) {
    buf := make([]byte, bufsize)
    // Read the incoming connection into the buffer.
    len, err := conn.Read(buf)
    fmt.Println("network.Read len : " , len)
    receive := string(buf[:len])
    return receive, err
}

func ConnectTo(addr string)(net.Conn, error) {
  conn, err := net.Dial("tcp", addr)
  if err != nil {
    fmt.Println("Error dialing:", err.Error())
    panic(err)
  }
  return conn, err
}

