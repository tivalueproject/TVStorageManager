package rpc

import (
	//"encoding/json"
	"fmt"
	"TVStorageManager/network"
  //"TVStorageManager/logic"
)


type SocketConfig struct {
  rpcAddress string
  tvRpcAddress string
}

// DefaultSocketConfig.
var DefaultSocketConfig = SocketConfig{
  rpcAddress:  "127.0.0.1:60000",
  tvRpcAddress : "127.0.0.1:63695",
}

func StartRpcServer() {

  InitRpcInterFaces()


  listener, err := network.Listen(DefaultSocketConfig.rpcAddress)
  if err != nil {
      fmt.Println("Error listening:", err.Error())
      panic(err)
  }

  // Close the listener when the application closes.
  defer listener.Close()
  fmt.Println(DefaultSocketConfig.rpcAddress)

  for {
    // Listen for an incoming connection.
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("Error accepting: ", err.Error())
      panic(err)
    } else { 

      // Connect to tv
      connTv, err := network.ConnectTo(DefaultSocketConfig.tvRpcAddress)

      requests, err := network.Read(conn)
      for err == nil {
        fmt.Println("Request from QT: " + requests)

        // ParseAndVerify
        parsedRequest, method := ParseVerify(requests)

        // Process by logic layer 
        response := ProcessRpcRequest(parsedRequest, method, connTv )

        // Send a response back to person contacting us.
        conn.Write( []byte( response))

        requests, err = network.Read(conn)
      }
    }
  }
}