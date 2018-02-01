package rpc

import (
	"net"
	"fmt"
	"TVStorageManager/json"
)
func ProcessRpcRequest(request *json.Json, method string, conn net.Conn)(response string) {

	if rpcInterface, ok := rpcInterFaceTable[method]; ok {
		fmt.Println(rpcInterface)
		response = rpcInterface.Func(request, conn)
	} else {
		fmt.Println("rpc method \"", method , "\" Not Found")
	}

	return
}

