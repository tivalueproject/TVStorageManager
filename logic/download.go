package logic

import (
	"net"
	"TVStorageManager/network"
)

func ListConfirmSavedFile(jsonStr string, conn net.Conn) (response string) {
	response = network.CallRpc(jsonStr, conn)
	return
 }

//cat file to local
func DownloadFileToLocal(hash string, filepath string, conn net.Conn) (response string) {
	err := Download(hash, filepath)
	if err != nil {
		panic(err)
	}

	return "200"
}

