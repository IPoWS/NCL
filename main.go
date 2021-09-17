package main

import (
	"fmt"
	"os"
	"time"

	"github.com/IPoWS/node-core/ip64"
	"github.com/IPoWS/node-core/link"
	"github.com/IPoWS/node-core/upper"
	"github.com/sirupsen/logrus"
)

type txtservice struct {
}

func (s *txtservice) Handle(srcport uint16, destport uint16, data *[]byte) {
	fmt.Println(data)
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugln("[NCL] start.")
	link.InitEntry("ws://"+os.Args[1]+"/nps", os.Args[2], os.Args[3], 0xffff_ffff_0000_0000)
	logrus.Debugln("[NCL] init ent succ.")
	err := link.Register()
	if err == nil {
		logrus.Debugln("[NCL] reg succ.")
		err = link.ListenAccess()
		if err == nil {
			logrus.Debugln("[NCL] listen succ.")
			ts := upper.Service(new(txtservice))
			upper.Register(1, &ts)
			logrus.Debugln("[NCL] service registered.")
			time.Sleep(time.Second)
			logrus.Debugln("[NCL] enter loop.")
			for {
				var i uint64
				s := make([]byte, 512)
				fmt.Print("Enter ip:")
				fmt.Scanf("%x", &i)
				fmt.Print("Enter msg:")
				fmt.Scanln(&s)
				_, err = link.Send(i, &s, ip64.DataType, 1, 1)
				if err != nil {
					logrus.Errorln("[NCL] send err:", err)
				} else {
					logrus.Infof("[NCL] send to %x succ.", i)
				}
			}
		} else {
			logrus.Errorln("[NCL] listen access err:", err)
		}
	}
}
