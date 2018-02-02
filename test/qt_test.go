package test

import (
	"fmt"
	"testing"
	"TVStorageManager/network"
)

// storage
// zxcmnbv65757TV8fvkrLJVDPhVB2AWAykJtZuG4UtCFLBDxV1SvNUWBfqduAjcDB
// func TestListFileStoreRequest(t * testing.T) {
// 	const login = `{"jsonrpc":"2.0","id":1,"method":"login","params":["admin","btc123"]}`
// 	const open = `{"jsonrpc":"2.0","id":1,"method":"open","params":["wallet"]}`
// 	const unlock = `{"jsonrpc":"2.0","id":1,"method":"unlock","params":[999999, "kg.com123"]}`
// 	const request = `{"jsonrpc":"2.0","id":2,"method":"ListFileStoreRequest","params":["1234567hashsbnssTV8fvkrLJVDPhVB2AWAykJtZuG4UtCFLBDxV1SvNUWBfqduAjcDB"]}`
// 	connTv, err := network.ConnectTo("127.0.0.1:60000")
// 	if err == nil {
// 		fmt.Fprintf(connTv, login)
// 		loginRes, _ := network.Read(connTv)
// 		if loginRes != "" {
// 			fmt.Println(loginRes)
// 		}
// 		fmt.Fprintf(connTv, open)
// 		openRes, _ := network.Read(connTv)
// 		if openRes != "" {
// 			fmt.Println(openRes)
// 		}
// 		fmt.Fprintf(connTv, unlock)
// 		unlockRes, _ := network.Read(connTv)
// 		if unlockRes != "" {
// 			fmt.Println(unlockRes)
// 		}
// 		fmt.Fprintf(connTv, request)
// 		response, _ := network.Read(connTv)
// 		if response != "" {
// 			fmt.Println(response)
// 		}
// 	}
// }


// func TestDeclareUploadFile(t * testing.T) {
// 	const login = `{"jsonrpc":"2.0","id":1,"method":"login","params":["admin","btc123"]}`
// 	const open = `{"jsonrpc":"2.0","id":1,"method":"open","params":["storage"]}`
// 	const unlock = `{"jsonrpc":"2.0","id":1,"method":"unlock","params":[999999, "kg.com123"]}`
// 	const request = `{"jsonrpc":"2.0","id":2,"method":"DeclareUploadFile","params":["lyy","123sbn","777","123testSQ","1234567hashsbnss;1234567hashsbnss,777;","TV","10","2","1","2","node4","1"]}`
// 	connTv, err := network.ConnectTo("127.0.0.1:60000")
// 	if err == nil {
// 		fmt.Fprintf(connTv, login)
// 		loginRes, _ := network.Read(connTv)
// 		if loginRes != "" {
// 			fmt.Println(loginRes)
// 		}
// 		fmt.Fprintf(connTv, open)
// 		openRes, _ := network.Read(connTv)
// 		if openRes != "" {
// 			fmt.Println(openRes)
// 		}
// 		fmt.Fprintf(connTv, unlock)
// 		unlockRes, _ := network.Read(connTv)
// 		if unlockRes != "" {
// 			fmt.Println(unlockRes)
// 		}
// 		fmt.Fprintf(connTv, request)
// 		response, _ := network.Read(connTv)
// 		if response != "" {
// 			fmt.Println(response)
// 		}
// 	}
// }


// func TestListUploadDeclaration(t * testing.T) {
// 	const login = `{"jsonrpc":"2.0","id":1,"method":"login","params":["admin","btc123"]}`
// 	const open = `{"jsonrpc":"2.0","id":1,"method":"open","params":["wallet"]}`
// 	const unlock = `{"jsonrpc":"2.0","id":1,"method":"unlock","params":[999999, "kg.com123"]}`
// 	const request = `{"jsonrpc":"2.0","id":2,"method":"ListUploadDeclaration","params":[]}`
// 	connTv, err := network.ConnectTo("127.0.0.1:60000")
// 	if err == nil {
// 		fmt.Fprintf(connTv, login)
// 		loginRes, _ := network.Read(connTv)
// 		if loginRes != "" {
// 			fmt.Println(loginRes)
// 		}
// 		fmt.Fprintf(connTv, open)
// 		openRes, _ := network.Read(connTv)
// 		if openRes != "" {
// 			fmt.Println(openRes)
// 		}
// 		fmt.Fprintf(connTv, unlock)
// 		unlockRes, _ := network.Read(connTv)
// 		if unlockRes != "" {
// 			fmt.Println(unlockRes)
// 		}
// 		fmt.Fprintf(connTv, request)
// 		response, _ := network.Read(connTv)
// 		if response != "" {
// 			fmt.Println(response)
// 		}
// 	}
// }

// func TestDeclarePieceSaved(t * testing.T) {
// 	const login = `{"jsonrpc":"2.0","id":1,"method":"login","params":["admin","btc123"]}`
// 	const open = `{"jsonrpc":"2.0","id":1,"method":"open","params":["wallet"]}`
// 	const unlock = `{"jsonrpc":"2.0","id":1,"method":"unlock","params":[999999, "kg.com123"]}`
// 	const request = `{"jsonrpc":"2.0","id":2,"method":"DeclarePieceSaved","params":["1234567hashsbnssTV8fvkrLJVDPhVB2AWAykJtZuG4UtCFLBDxV1SvNUWBfqduAjcDB","1234567hashsbnss","lyy2","node7"]}`
// 	connTv, err := network.ConnectTo("127.0.0.1:60000")
// 	if err == nil {
// 		fmt.Fprintf(connTv, login)
// 		loginRes, _ := network.Read(connTv)
// 		if loginRes != "" {
// 			fmt.Println(loginRes)
// 		}
// 		fmt.Fprintf(connTv, open)
// 		openRes, _ := network.Read(connTv)
// 		if openRes != "" {
// 			fmt.Println(openRes)
// 		}
// 		fmt.Fprintf(connTv, unlock)
// 		unlockRes, _ := network.Read(connTv)
// 		if unlockRes != "" {
// 			fmt.Println(unlockRes)
// 		}
// 		fmt.Fprintf(connTv, request)
// 		response, _ := network.Read(connTv)
// 		if response != "" {
// 			fmt.Println(response)
// 		}
// 	}
// }


// func TestConfirmFileSaved(t * testing.T) {
// 	const login = `{"jsonrpc":"2.0","id":1,"method":"login","params":["admin","btc123"]}`
// 	const open = `{"jsonrpc":"2.0","id":1,"method":"open","params":["storage"]}`
// 	const unlock = `{"jsonrpc":"2.0","id":1,"method":"unlock","params":[999999, "kg.com123"]}`
// 	const request = `{"jsonrpc":"2.0","id":2,"method":"ConfirmFileSaved","params":["lyy","1234567hashsbnssTV8fvkrLJVDPhVB2AWAykJtZuG4UtCFLBDxV1SvNUWBfqduAjcDB","1234567hashsbnss","TV8mCM3p9wXr4wDAz93X2ceSkPmN5rCtZqZ3Df1c33562sRbceug","1"]}`
// 	connTv, err := network.ConnectTo("127.0.0.1:60000")
// 	if err == nil {
// 		fmt.Fprintf(connTv, login)
// 		loginRes, _ := network.Read(connTv)
// 		if loginRes != "" {
// 			fmt.Println(loginRes)
// 		}
// 		fmt.Fprintf(connTv, open)
// 		openRes, _ := network.Read(connTv)
// 		if openRes != "" {
// 			fmt.Println(openRes)
// 		}
// 		fmt.Fprintf(connTv, unlock)
// 		unlockRes, _ := network.Read(connTv)
// 		if unlockRes != "" {
// 			fmt.Println(unlockRes)
// 		}
// 		fmt.Fprintf(connTv, request)
// 		response, _ := network.Read(connTv)
// 		if response != "" {
// 			fmt.Println(response)
// 		}
// 	}
// }

// func TestListConfirmFileSaved(t * testing.T) {
// 	const login = `{"jsonrpc":"2.0","id":1,"method":"login","params":["admin","btc123"]}`
// 	const open = `{"jsonrpc":"2.0","id":1,"method":"open","params":["storage"]}`
// 	const unlock = `{"jsonrpc":"2.0","id":1,"method":"unlock","params":[999999, "kg.com123"]}`
// 	const request = `{"jsonrpc":"2.0","id":2,"method":"ListConfirmFileSaved","params":[]}`
// 	connTv, err := network.ConnectTo("127.0.0.1:60000")
// 	if err == nil {
// 		fmt.Fprintf(connTv, login)
// 		loginRes, _ := network.Read(connTv)
// 		if loginRes != "" {
// 			fmt.Println(loginRes)
// 		}
// 		fmt.Fprintf(connTv, open)
// 		openRes, _ := network.Read(connTv)
// 		if openRes != "" {
// 			fmt.Println(openRes)
// 		}
// 		fmt.Fprintf(connTv, unlock)
// 		unlockRes, _ := network.Read(connTv)
// 		if unlockRes != "" {
// 			fmt.Println(unlockRes)
// 		}
// 		fmt.Fprintf(connTv, request)
// 		response, _ := network.Read(connTv)
// 		if response != "" {
// 			fmt.Println(response)
// 		}
// 	}
// }

// func TestUploadFileToIPFS(t * testing.T) {
// 	const request = `{"jsonrpc":"2.0","id":2,"method":"UploadFileToIPFS","params":["F:\\go\\src\\TVStorageManager\\README.md"]}`
// 	connTv, err := network.ConnectTo("127.0.0.1:60000")
// 	if err == nil {
// 		fmt.Fprintf(connTv, request)
// 		response, _ := network.Read(connTv)
// 		if response != "" {
// 			fmt.Println(response)
// 		}
// 	}
// }

// func TestPinAddFileToLocal(t * testing.T) {
// 	const request = `{"jsonrpc":"2.0","id":2,"method":"PinAddFileToLocal","params":["QmQy7RRsJE6ueskwtTwzZNpPhSBksoZTBBqAxwhwjihNv7"]}`
// 	connTv, err := network.ConnectTo("127.0.0.1:60000")
// 	if err == nil {
// 		fmt.Fprintf(connTv, request)
// 		response, _ := network.Read(connTv)
// 		if response != "" {
// 			fmt.Println(response)
// 		}
// 	}
// }

func TestDownloadFileFromIPFS(t * testing.T) {
	const request = `{"jsonrpc":"2.0","id":2,"method":"DownloadFileFromIPFS","params":["QmQy7RRsJE6ueskwtTwzZNpPhSBksoZTBBqAxwhwjihNv7"]}`
	connTv, err := network.ConnectTo("127.0.0.1:60000")
	if err == nil {
		fmt.Fprintf(connTv, request)
		response, _ := network.Read(connTv)
		if response != "" {
			fmt.Println(response)
		}
	}
}