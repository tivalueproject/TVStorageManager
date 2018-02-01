package rpc

import (
	"fmt"
	"TVStorageManager/json"
	"net"
)


// passThroughTable
type RpcInterFace struct {
	Method      	string
	Func 			func(request *json.Json, tvConn net.Conn)(response string)
	IsPassThrough	bool
}

var rpcInterFaceTable map[string]RpcInterFace
func InitRpcInterFaces() {

	rpcInterFaceTable = make(map[string]RpcInterFace)

	rpcInterFaceTable["UploadFileToIPFS"] = RpcInterFace{"UploadFileToIPFS", UploadFileToIPFS, false}
	rpcInterFaceTable["DeclareUploadFile"] = RpcInterFace{"DeclareUploadFile", DeclareUploadFile, false}

	rpcInterFaceTable["login"] = RpcInterFace{"login", RpcPassThrough, true}
	rpcInterFaceTable["login2"] = RpcInterFace{"login2", RpcPassThrough, true}
}


func RpcPassThrough(request *json.Json, tvConn net.Conn)(response string) {
	
 	if request != nil {
		//id := request.Get("id")
		method, _ := request.Get("method").String()
		params, _ := request.Get("params").Array()
	    fmt.Println("RpcPassThrough called ", method , params[0], params[1])
	    requestStr := json.GenerateJsonString(method, params)
	    response = CallRpc(requestStr, tvConn)
	    fmt.Println(response)
	}
	return
}

func UploadFileToIPFS(request *json.Json, tvConn net.Conn)(response string) {

	fmt.Println("UploadFileToIPFS called ")
	return
}

func DeclareUploadFile(request *json.Json, tvConn net.Conn)(response string) {

	fmt.Println("DeclareUploadFile called ")
	return
}
