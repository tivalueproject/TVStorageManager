package main

import (
  "TVStorageManager/rpc"
  "time"
)

func main() {

	// boot TV
	
	// boot ipfs

	// boot rpc service
	go rpc.StartRpcServer()

  	for {
		time.Sleep(1000)
  	}

}