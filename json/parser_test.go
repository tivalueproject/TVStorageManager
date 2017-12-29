package json

import (
	"fmt"
	"testing"
	"encoding/json"
	"TVStorageManager/common"
)

func TestJSONParser(t * testing.T) {
	const jsonStream = `{"jsonrpc":"2.0","id":1,"method":"login","params":["username","password"]}`
	result, _ := JsonParser([]byte(jsonStream))
	if result != nil {
		str, _ := result.Get("method").String()
		fmt.Println("string: ", str)
		arr, _ := result.Get("params").Array()
		fmt.Println("array: ", arr[0])
		fmt.Println("interface{}: ", result.Interface())
	}
}

func TestGenerateJSON(t * testing.T) {
	const jsonStream = `{"jsonrpc":"2.0","id":1,"method":"login","params":["username","password"]}`
	result, _ := JsonParser([]byte(jsonStream))
	if result != nil {
		arr, _ := result.Get("params").Array()
		var jsonStr = GenerateJsonString("login", arr)
		fmt.Println("generate json string: ", jsonStr)
	}
}

func TestPaserJsonString(t * testing.T) {
	const jsonStream = `{"id":1,"result":[{"id":{"file_id":"658e69ec63055618512e4ddadb3742b325f85540e31841f513aa9a70f9a16919","uploader":"TV7aGRJEjKdLJ9TCMyw8E9PQCTjVHGxvA8dekAvXzbyRKkdkK46C"},"pieces":[{"pieceid":"Qmc7cK4vaVc1Pet7DdsVkeWLizzVaFLTbkgYoS6nSfutmw","piece_size":19718,"price":"10.000000000000"}],"num_of_copys":1,"node_id":"/ip4/172.16.0.206/tcp/4001/ipfs/QmduzrWsgH4WYTXY4zUPZPhBSfeVLt5ifN4GiJmMMmVVev","filename":"raw.hpp","description":"myfiledescription"},{"id":{"file_id":"4b838aacdc0584c46337f79a7ded87cf","uploader":"TV7aGRJEjKdLJ9TCMyw8E9PQCTjVHGxvA8dekAvXzbyRKkdkK46C"},"pieces":[{"pieceid":"QmY1jQzEts5wdTrGN1kAtMttXh7MTAoAQJCeYKeXKQKmxr","piece_size":4361,"price":"15.000000000000"}],"num_of_copys":1,"node_id":"/ip4/127.0.0.1/tcp/4001/ipfs/QmduzrWsgH4WYTXY4zUPZPhBSfeVLt5ifN4GiJmMMmVVev","filename":"local_rpc3.py","description":"mydes"}]}`
	fmt.Println("jsonString: ", jsonStream)
	res, _ := JsonParser([]byte(jsonStream))
	if res != nil {
		arr, _ := res.Get("result").Array()
		//fmt.Println(arr)
		req, _ := json.Marshal(arr[0])
		var reqInfo common.RequestInfo
		json.Unmarshal(req, &reqInfo)
		fmt.Println("fileId: ", reqInfo.Id.FileId)
		fmt.Println("copies: ", reqInfo.Copies)
	}
}