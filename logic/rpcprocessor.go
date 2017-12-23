package logic

import (
	"net"
	"fmt"
)
func ProcessRpcRequest(request interface{}, isPassThrough bool, conn net.Conn)(response string) {

	if isPassThrough {
		response = CallRpc(request.(string), conn)
        fmt.Println("Response from server: " + response)
	} else {
		//UploadFile("F:\\qt-opensource-windows-x86-msvc2013_64-5.6.3.exe", 1, 1, conn)
	}

	return
}