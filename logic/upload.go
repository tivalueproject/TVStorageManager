package logic

import (
	// "strconv"
	"os"
	"net"
	myJson "TVStorageManager/json"
	"fmt"
	"TVStorageManager/network"
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
func UploadFileToIPFS(file_name string, conn net.Conn)(res interface{}, err error) {
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
	// res = "{\"id\":\"1\",\"file_name\":\"" + name + "\",\"file_hash\":\"" + hash + "\",\"file_size\":\"" + size + "\",\"status_code\":\"" + strconv.Itoa(stat_code) + "\"}"
	var s = make(map[string]interface{})
	s["id"] = "1"
	s["file_name"] = name
	s["file_hash"] = hash
	s["file_size"] = size
	s["status_code"] = stat_code
	res = s
	return res, nil
}


//pin file to local
func PinAddFileToLocal(hash string, conn net.Conn)(res interface{}, err error) {
	resp, err := PinAdd(hash)
	result, _ := myJson.JsonParser([]byte(resp))
	if err != nil {
    	panic(err)
	}
	obj, _ := result.Get("Pins").String()
	// res = "{\"id\":\"1\",\"pins\":\"" + obj + "\""+ ",\"status_code\":\"" + strconv.Itoa(stat_code) + "\"}"
	if obj != "" {
		var s = make(map[string]interface{})
		s["Pins"] = obj
		s["status_code"] = 200
		res = s
	}
	return res, nil
}

//list upload requests
func ListUploadedRequests(conn net.Conn)(response string) {
	// strjson := "{\"jsonrpc\":\"2.0\",\"method\":\"list_upload_requests\",\"params\":\"[]\",\"id\":\"1\"}"
	strjson := myJson.GenerateJsonString("list_upload_requests","[]")
	resp := network.CallRpc(strjson, conn)
	return resp
}

 //get local node id
func GetLocalNodeId()(res string, err error) {
	resp, err := Id()

	result, _ := myJson.JsonParser([]byte(resp))
	if err != nil {
    	panic(err)
	}
	nodeId, _ := result.Get("ID").String()
	res = "{\"id\":\"" + nodeId + "\"}"
	return res, nil
}