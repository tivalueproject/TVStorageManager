package rpc

import (
	//"encoding/json"
	"fmt"
	"TVStorageManager/json"
)


// passThroughTable
type PassThroughTable struct {


}

func ParseVerify(json string) (*json.Json, string) {

	var jsonRequest = ParseJson(json)
	var methodName string = ""
	if jsonRequest != nil {

		//jsonrpc, _ := jsonRequest.Get("jsonrpc").String()
		id := jsonRequest.Get("id")
		method, _ := jsonRequest.Get("method").String()
		params, _ := jsonRequest.Get("params").Array()
		// if jsonrpc!="2.0" {
		// 	fmt.Println("Invalid JsonRpc: ", json, "field \"jsonrpc\"  be specified and must be \"2.0\"!")
		// 	return nil, ""
		// } else 
		if id==nil {
			fmt.Println("Invalid JsonRpc: ", json, "field \"id\" must be specified!")
			return nil, ""
		} else if method=="" {
			fmt.Println("Invalid JsonRpc: ", json, "field \"method\" must be specified!")
			return nil, ""
		} else if params==nil {
			fmt.Println("Invalid JsonRpc: ", json, "field \"params\" must be specified, and must be an array!")
			return nil, ""
		}
		methodName = method


	}

	return jsonRequest, methodName
}

func ParseJson(jsonstr string) *json.Json {
	result, err := json.JsonParser([]byte(jsonstr))
	if result != nil {
		return result
	} else {
		fmt.Println("Error ParseJson: ", err.Error())
		return nil
	}
}
