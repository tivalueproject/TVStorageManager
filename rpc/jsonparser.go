package rpc

import (
	//"encoding/json"
	"fmt"
	"TVStorageManager/json"
)


// passThroughTable
type PassThroughTable struct {


}

func Parse(json string) (interface{}, bool) {
	if true/* in the passThroughTable */ {
	    return json, true
	} else {
	   return ParseJson(json), false
	}
}

func ParseJson(jsonstr string) interface{} {
	result, err := json.JsonParser([]byte(jsonstr))
	if result != nil {
		return result
	} else {
		fmt.Println("Error ParseJson: ", err.Error())
		return nil
	}
}
