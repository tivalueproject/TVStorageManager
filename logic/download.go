package logic

import (
	"net"
	"TVStorageManager/network"	
)

func ListSavedFile(jsonStr string, conn net.Conn) (resp string) {
	response := network.CallRpc(jsonStr, conn)
	return response
 }

// func GenerateDownloadValidation(params []string, conn net.Conn)(response string) {
// 	var strjson string
// 	if params != nil {
// 		strjson := TVJson.GenerateJsonString("generate_download_validation", params)
// 	}
 	
// 	resp := CallRpc(strjson, conn)
// 	return resp
// }

