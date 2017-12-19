package json

import (
	"fmt"
	"testing"
)

func TestParse(t * testing.T) {
	const jsonStream = `{"jsonrpc":"2.0","id":1,"method":"login","params":["username","password"]}`

	result, str, _ := parse(jsonStream)
	fmt.Println("i store string: ", str)
	fmt.Println("map: ", result)
}

func TestParseOne(t * testing.T) {
	const jsonStream = `{"jsonrpc":"2.0","id":1,"method":"login","params":["username","password"]}`

	result, _ := parseOne(jsonStream)
	switch t := result.(type) {
	case string:
		fmt.Println("i store string: ", t)
	default:
		fmt.Println("map: ", t)
	}
}

func TestParseObjectWithRightKey(t * testing.T) {
	const jsonStream = `{"jsonrpc":"2.0","id":1,"method":"info","params":[]}`
	result, _ := parseObject(jsonStream, "method")
	if result != nil {
		t.Log("method", result)
	}
}

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