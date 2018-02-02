package test

import (
	"fmt"
	"testing"
	"TVStorageManager/network"
)

func TestListFileStoreRequest(t * testing.T) {
	const login = `{"jsonrpc":"2.0","id":1,"method":"login","params":["admin","btc123"]}`
	const open = `{"jsonrpc":"2.0","id":1,"method":"open","params":["storage"]}`
	const unlock = `{"jsonrpc":"2.0","id":1,"method":"unlock","params":[999999, "kg.com123"]}`
	const request = `{"jsonrpc":"2.0","id":2,"method":"ListFileStoreRequest","params":["zxcmnbv65757TV8fvkrLJVDPhVB2AWAykJtZuG4UtCFLBDxV1SvNUWBfqduAjcDB"]}`
	connTv, err := network.ConnectTo("127.0.0.1:60000")
	if err == nil {
		fmt.Fprintf(connTv, login)
		loginRes, _ := network.Read(connTv)
		if loginRes != "" {
			fmt.Println(loginRes)
		}
		fmt.Fprintf(connTv, open)
		openRes, _ := network.Read(connTv)
		if openRes != "" {
			fmt.Println(openRes)
		}
		fmt.Fprintf(connTv, unlock)
		unlockRes, _ := network.Read(connTv)
		if unlockRes != "" {
			fmt.Println(unlockRes)
		}
		fmt.Fprintf(connTv, request)
		response, _ := network.Read(connTv)
		if response != "" {
			fmt.Println(response)
		}
	}
}