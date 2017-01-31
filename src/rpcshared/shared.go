package rpcshared

import (
	"regexp"
    "fmt"
)

type GrepTool struct {
	NumberRequests int
	RequestHistory []float64
}

type Args struct {
	Data   []byte
	FileID string
}

func (t *GrepTool) Execute(args *Args, reply *string) error {
    fmt.Println("I got data of length: ", len(args.Data))
   regex, err := regexp.Compile("(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	if err != nil {
return err
}
	        if regex.Match(args.Data) {
			*reply = "true"
		} else {
			*reply = "false"
        }
return err
}
