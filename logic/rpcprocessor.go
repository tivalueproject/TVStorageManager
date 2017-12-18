package logic

import (
	"net"
	"fmt"
)
func ProcessRpcRequest(request interface{}, isPassThrough bool, conn net.Conn)(response string) {

	if isPassThrough {
		response = CallRpc(request.(string),conn)
        fmt.Println("Response from server: " + response)
	} else {

	}

	return
}