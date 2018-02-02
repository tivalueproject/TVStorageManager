package logic

import (
	"strconv"
	"os"
	"net"
	myjson "TVStorageManager/json"	
	//"encoding/json"
	"fmt"
	"TVStorageManager/rpc"
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
	result, _ := myjson.JsonParser([]byte(resp))
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
	tv_res := rpc.CallRpc(json_prc_str, conn)
	return tv_res, nil
}

//list uploaded declaration
func ListUploadDeclaration(json_rpc_str string, conn net.Conn)(res string, err error) {
	tv_res := rpc.CallRpc(json_rpc_str, conn)
	return tv_res, nil
}

//pin file to local
func PinAddFileToLocal(hash string, conn net.Conn)(res string, err error) {
	resp, stat_code, err := PinAdd(hash)
	if 200 != stat_code {
		panic(err)
	}
	result, _ := myjson.JsonParser([]byte(resp))
	if err != nil {
    	panic(err)
	}
	obj, _ := result.Get("Pins").String()
	res = "{\"id\":\"1\",\"pins\":\"" + obj + "\""+ ",\"status_code\":\"" + strconv.Itoa(stat_code) + "\"}"
	return res, nil
}

//declare piece saveds
func DeclarePieceSaved(json_rpc_str string, conn net.Conn) (res string, err error) {
	tv_res := rpc.CallRpc(json_rpc_str, conn)
	return tv_res, nil
}


//list upload requests
func ListUploadedRequests(conn net.Conn)(response string) {
	strjson := "{\"jsonrpc\":\"2.0\",\"method\":\"list_upload_requests\",\"params\":\"[]\",\"id\":\"1\"}"
	resp := rpc.CallRpc(strjson, conn)
	return resp
}

//save piece
func SavePiece(params []string, conn net.Conn)(resp string) {
	strjson := ""
	return rpc.CallRpc(strjson, conn)
}

//list store request
 func ListStoreRequest(conn net.Conn) {

 }

 //allow save
 func AllowSave(conn net.Conn) {

 }

 //list save declaration
 func ListSaveDeclaration(conn net.Conn) {

 }

 //confirm piece
 func ConfirmPiece(conn net.Conn) {

 }

 