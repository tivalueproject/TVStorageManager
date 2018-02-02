package logic

import (
	"strconv"
	"os"
	"net"
	myJson "TVStorageManager/json"	
	"encoding/json"
	"fmt"
	"TVStorageManager/network"
	"TVStorageManager/common"
)


const (
	MinPieceSize int64 = 4096
	BuffSize int64 = 2 << 10
)

//copy file
func CopyFile(fh *os.File, len int64, filename string) {
	buff := make([]byte, BuffSize)
	fhp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	for {
		readlen, _ := fh.Read(buff)
		fhp.Write(buff)
		len -= int64(readlen)
		if len <= 0 {
			break
		}
	}

}

//slice file
func Slice(filename string, pieces int) {
	fileinfo, _ := os.Stat(filename)
	filesize := fileinfo.Size()
	piecesize := filesize / int64(pieces)
	
	if filesize <= MinPieceSize {
		pieces = 1
	} else if piecesize < MinPieceSize {
		pieces = 2
	}

	f, _ := os.Open(filename)

	for i := 0; i < pieces; i++ {
		piecename := fmt.Sprintf("%s_%d", fileinfo.Name(), i)
		if i == pieces - 1 {
			CopyFile(f, filesize - piecesize * int64((pieces - 1)), piecename)
		} 
		CopyFile(f, piecesize, piecename)
	}
}

//upload file to local ipfs
func UploadFileToIPFS(file_name string, conn net.Conn)(res string, err error) {
	resp, stat_code, err := Upload(file_name)
	if 200 != stat_code {
		panic(err)
	}
	result, _ := myJson.JsonParser([]byte(resp))
	if err != nil {
    	panic(err)
    }
	name, _ := result.Get("Name").String()
	hash, _ := result.Get("Hash").String()
	size, _ := result.Get("Size").String()
	res = "{\"id\":\"1\",\"file_name\":\"" + name + "\",\"file_hash\":\"" + hash + "\",\"file_size\":\"" + size + "\",status_code:\"" + strconv.Itoa(stat_code) + "\"}"
	return res, nil
}

//
func DeclareUploadFile(json_prc_str string, conn net.Conn) (res string, err error) {
	tv_res := network.CallRpc(json_prc_str, conn)
	return tv_res, nil
}

//list uploaded declaration
func ListUploadDeclaration(json_rpc_str string, conn net.Conn)(res string, err error) {
	tv_res := network.CallRpc(json_rpc_str, conn)
	return tv_res, nil
}

//pin file to local
func PinAddFileToLocal(hash string, conn net.Conn)(res string, err error) {
	resp, stat_code, err := PinAdd(hash)
	if 200 != stat_code {
		panic(err)
	}
	result, _ := myJson.JsonParser([]byte(resp))
	if err != nil {
    	panic(err)
	}
	obj, _ := result.Get("Pins").String()
	res = "{\"id\":\"1\",\"pins\":\"" + obj + "\""+ ",\"status_code\":\"" + strconv.Itoa(stat_code) + "\"}"
	return res, nil
}

//declare piece saveds
func DeclarePieceSaved(json_rpc_str string, conn net.Conn) (res string, err error) {
	tv_res := network.CallRpc(json_rpc_str, conn)
	return tv_res, nil
}


//list upload requests
func ListUploadedRequests(conn net.Conn)(response string) {
	strjson := "{\"jsonrpc\":\"2.0\",\"method\":\"list_upload_requests\",\"params\":\"[]\",\"id\":\"1\"}"
	resp := network.CallRpc(strjson, conn)
	return resp
}

//save piece
func SavePiece(params []string, conn net.Conn)(resp string) {
	strjson := ""
	return network.CallRpc(strjson, conn)
}

//list store request
func ListStoreRequest(jsonStr string, conn net.Conn) (resp string) {
	var resJson string
	fmt.Println("call before: ", jsonStr)
	response := network.CallRpc(jsonStr, conn)
	fmt.Println("call after: ", response)
	res, _ := myJson.JsonParser([]byte(response))
	if res != nil {
		resArr, _ := res.Get("result").Array()
		for i := 0; i < len(resArr); i++ {
			resMapStr, _ := json.Marshal(resArr[i])
			var storeRes common.StoreFileResult
			json.Unmarshal(resMapStr, &storeRes)
			fmt.Println("fileId: ", storeRes.FileId.FileId)
			nodes := storeRes.Nodes
			if (nodes != nil) {
				for j := 0; j < len(nodes); j++ {
					nodesStr, _ := json.Marshal(nodes[i])
					var storeNode common.StoreNode
					json.Unmarshal(nodesStr, &storeNode)
					var s = make(map[string]interface{})
					s["node"] = storeNode.NodeId
					s["key"] = storeNode.NodeKey
					result, _ := json.Marshal(s)
					resJson = string(result)
					fmt.Println("result: ", resJson)
				}
			}
		}
	}
	return resJson
 }

 //confirm piece
 func ConfirmPieceSaved(jsonStr string, conn net.Conn) (resp string) {
	var resJson string
	response := network.CallRpc(jsonStr, conn)
	res, _ := myJson.JsonParser([]byte(response))
	if res != nil {
		var s = make(map[string]interface{})
		s["result"] = true
		result, _ := json.Marshal(s)
		resJson = string(result)
		fmt.Println("result: ", resJson)
	}
	return resJson
 }

 