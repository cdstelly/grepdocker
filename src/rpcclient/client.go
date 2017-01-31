package main

import (
	"io/ioutil"
	"fmt"
	"net/rpc"
//	"net/http"
	"log"
	"rpcshared"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:5553")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	filepath := "test.file"
	fileData, err:= ioutil.ReadFile(filepath)
	if err != nil { 
		log.Fatal("error reading file: ", err)
	}

	// Execute the task
	args := &rpcshared.Args{FileID: "test", Data: fileData}
	var reply string
	err = client.Call("GrepTool.Execute", args, &reply)
	if err != nil {
		log.Fatal("exif error:", err)
	}
	fmt.Printf("Result: %s\n", reply)

}
