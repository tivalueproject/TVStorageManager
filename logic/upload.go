package logic

import (
	//"strconv"
	"os"
	"net"
	//"TVStorageManager/json"	
	//json2 "encoding/json"
	"fmt"



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

// //upload file
// func UploadFile(filename string, piecesize int, pieces int,price float64, conn net.Conn) (response string) {	
// 	info, _ := os.Stat(filename)
// 	tvFileInfo := TVFileInfo{}
// 	tvFileInfo.FileName = info.Name()
// 	tvFileInfo.FileSize = info.Size()
// 	tvFileInfo.FileHash = "file_hash"
// 	tvFileInfo.Copies = 1
// 	tvFileInfo.Price = price
// 	tvFileInfo.Description = "my_description"
// 	tvFileInfo.Uploader = "TV3HtJtjsyZTQcAZyVZdrL79MMdx6zkfaSy"
// 	tvFileInfo.Contract = "CONAZ6B5WC2zybXq4UiEF57Lb9cvRNsY5Gjx"
// 	tvFileInfo.Pieces = make([]Piece, pieces)

// 	//slice file
// 	piecesize = int(info.Size())
// 	for i := 0; i < pieces; i++ {
// 		resp, err := Upload(filename)
// 		if err != nil {
// 			panic(err)
// 		}
// 		result, _ := json.JsonParser([]byte(resp))
// 		if err != nil {
// 			panic(err)
// 		}			
// 		hash, _ := result.Get("Hash").String()
// 		//size, _ := result.Get("Size").String()
// 		piece := Piece{}	
// 		piece.Id = hash
// 		piece.Size = int(info.Size())
// 		piece.Price = price / float64(pieces)
// 		tvFileInfo.Pieces[i] = piece
// 	}
// 	str, err := json2.Marshal(tvFileInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	str1 := "{\"jsonrpc\":\"2.0\",\"method\":\"store_file_to_network\",\"id\":1,\"params\":["
// 	str2 := "]}"
// 	str3 := str1  + string(str)[1:len(string(str))-1]+ str2
// 	fmt.Println(string(str3))

// 	return CallRpc(str3, conn)
// }

//list upload requests
func ListUploadedRequests(conn net.Conn)(response string) {
	strjson := "{\"jsonrpc\":\"2.0\",\"method\":\"list_upload_requests\",\"params\":\"[]\",\"id\":1}"
	resp := CallRpc(strjson, conn)
	return resp
}

//save piece
func SavePiece(params []string, conn net.Conn)(resp string) {
	strjson := ""
	return CallRpc(strjson, conn)
}

//list store request
 func ListStoreRequest(conn net.Conn) {

 }

 //allow save
 func AllowSave(conn net.Conn) {

 }

 //declare piece saved
 func DeclarePieceSaved(conn net.Conn) {

 }

 //list save declaration
 func ListSaveDeclaration(conn net.Conn) {

 }

 //confirm piece
 func ConfirmPiece(conn net.Conn) {

 }

 