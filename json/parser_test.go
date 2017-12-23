package json

import (
	"fmt"
	"testing"
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