package logic

import (
    "net"
    "fmt"
    "TVStorageManager/network"
)
func CallRpc(request string, conn net.Conn)(response string) {
    fmt.Fprintf(conn, request)
    response, _ = network.Read(conn)
    return response
}