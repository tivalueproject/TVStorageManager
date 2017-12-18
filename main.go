package main

import (
  "TVStorageManager/rpc"
)

func main() {

  go rpc.StartRpcServer()
  for {

  }

}