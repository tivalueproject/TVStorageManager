package main

import (
  "TVStorageManager/rpc"
  "TVStorageManager/test"
  "time"
  "fmt"
  "os/exec"
)

func main() {

	// boot TV
    tvCmd := exec.Command("Ti_Value.exe", "--daemon", "--data-dir", "C:\\Users\\Siegfried\\AppData\\Roaming\\TVWallet\\chaindata", "--rpcuser", "a", "--rpcpassword", "b", "--rpcendpoint", "127.0.0.1:63695")
    tvIn, _ := tvCmd.StdinPipe()
    tvOut, _ := tvCmd.StdoutPipe()
    tvCmd.Start()
    fmt.Printf("starting Ti_Value, [PID] %d running...\n", tvCmd.Process.Pid)
    tvIn.Close()
    tvOut.Close()

	// boot ipfs
    ipfsCmd := exec.Command("ipfs.exe", "daemon")
    ipfsIn, _ := ipfsCmd.StdinPipe()
    ipfsOut, _ := ipfsCmd.StdoutPipe()
    ipfsCmd.Start()
    fmt.Printf("starting ipfs, [PID] %d running...\n", ipfsCmd.Process.Pid)
    ipfsIn.Close()
    ipfsOut.Close()

	// boot rpc service
	go rpc.StartRpcServer()

	test.IpfsTest()

  	for {
		time.Sleep(1000)
  	}

}
