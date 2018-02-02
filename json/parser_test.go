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

func TestPaserStoreRequestJsonString(t * testing.T) {
	const jsonStream = `[{"file_id":{"file_id":"zxcmnbv65757","uploader":"TV8fvkrLJVDPhVB2AWAykJtZuG4UtCFLBDxV1SvNUWBfqduAjcDB"},"piece_id":"zxcmnbv65757","nodes":[{"node":"node_lyy2","key":"TV8mCM3p9wXr4wDAz93X2ceSkPmN5rCtZqZ3Df1c33562sRbceug"}]}]`
	//fmt.Println("jsonString: ", jsonStream)
	res, _ := JsonParser([]byte(jsonStream))
	if res != nil {
		resArr, _ := res.Array()
		for i := 0; i < len(resArr); i++ {
			resMapStr, _ := json.Marshal(resArr[i])
			var storeRes common.StoreFileResult
			json.Unmarshal(resMapStr, &storeRes)
			fmt.Println("fileId: ", storeRes.FileId.FileId)
			fmt.Println("nodes: ", storeRes.Nodes)
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
					fmt.Println("result: ", string(result))
				}
			}
		}
	}
}