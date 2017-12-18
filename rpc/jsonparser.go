package rpc

import (
	//"encoding/json"
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

func ParseJson(json string) interface{} {

	return nil
}
