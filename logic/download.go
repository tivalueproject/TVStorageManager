package logic

import (
	"net"
	"TVStorageManager/network"
	myJson "TVStorageManager/json"
	"encoding/json"
	"fmt"	
)

func ListConfirmSavedFile(jsonStr string, conn net.Conn) (response string) {
	response = network.CallRpc(jsonStr, conn)
	return
 }

//cat file to local
func DownloadFileToLocal(hash string, conn net.Conn)(resJson string) {
	resp, stat_code, err := Download(hash)
	if 200 != stat_code {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(stat_code)
	res, _ := myJson.JsonParser([]byte(resp))
	if res != nil {
		var s = make(map[string]interface{})
		s["status"] = stat_code
		s["data"] = resp
		result, _ := json.Marshal(s)
		resJson = string(result)
	}
	return 
}

