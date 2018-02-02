package network

import (
    "net"
    "fmt"
)
func CallRpc(request string, conn net.Conn)(response string) {
    fmt.Fprintf(conn, request)
    response, _ = Read(conn)
    return response
}