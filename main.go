package main

import (
	"fmt"
	"os"

	"github.com/IPoWS/node-core/ip64"
	"github.com/IPoWS/node-core/link"
	"github.com/IPoWS/node-core/upper"
)

type txtservice struct {
}

func (s *txtservice) Handle(srcport uint16, destport uint16, data *[]byte) {
	fmt.Println(data)
}

func main() {
	link.InitEntry("ws://"+os.Args[1]+"/nps", os.Args[2], os.Args[3], 0xffff_ffff_0000_0000)
	link.Register()
	ts := upper.Service(new(txtservice))
	upper.Register(1, &ts)
	for {
		var i uint64
		var s []byte
		fmt.Print("Enter ip:")
		fmt.Scanf("%x", &i)
		fmt.Print("Enter msg:")
		fmt.Scanln(&s)
		link.Send(i, &s, ip64.DataType, 1, 1)
	}
}
