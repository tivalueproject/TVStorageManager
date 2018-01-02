package logic

import (
	"net"
	TVJson "TVStorageManager/json"	
)

func ListSaveNodes(params []string, conn net.Conn)(response string) {
	var strjson string
	if params != nil {
		strjson := TVJson.GenerateJsonString("blockchain_get_file_save_node", params)
	}
 	
	resp := CallRpc(strjson, conn)
	return resp
}

func GenerateDownloadValidation(params []string, conn net.Conn)(response string) {
	var strjson string
	if params != nil {
		strjson := TVJson.GenerateJsonString("generate_download_validation", params)
	}
 	
	resp := CallRpc(strjson, conn)
	return resp
}

