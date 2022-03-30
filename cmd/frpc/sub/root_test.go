package sub

import (
	"github.com/fatedier/golib/crypto"
	"github.com/panglove/frp/pkg/util"
	"log"
	"testing"
	"time"
)
func TestRunClient2(t *testing.T) {
	go ConnectFrpServer()
	select {

	}
}
func TestRunClient(t *testing.T) {
	go ConnectFrpServer()
	select {

	}
}
func ProcessClientInfo(){


		time.Sleep(time.Second*5)
		log.Println("refresh info",time.Now().String(),NewService)
		if NewService!=nil {
			RunClientReload("127.0.0.1", 20001, 22, 6000, "mine5name","HELLO2222")

			//time.Sleep(time.Second*1111)
			//ips := util.ExternalIP()
			//ipsStr := ""
			//if ips!=nil && len(ips)>0 {
			//	for _,v :=range  ips {
			//		ipsStr =ipsStr+ v+"|"
			//	}
			//	ipsStr = ipsStr[:len(ipsStr)-1]
			//}
			//log.Println("refresh info",time.Now().String())
			//RunClient("127.0.0.1", 20001, 22, 6000, "mine3name",ipsStr+"|"+time.Now().String())
		}

}
func ConnectFrpServer()  {
	crypto.DefaultSalt = "frp"
	ips := util.ExternalIP()
	ipsStr := ""
	if ips!=nil && len(ips)>0 {
		for _,v :=range  ips {
			ipsStr =ipsStr+ v+"|"
		}
		ipsStr = ipsStr[:len(ipsStr)-1]
	}

	//go ProcessClientInfo()
	for {
		err := RunClient("127.0.0.1", 7000, 22, 6000, "mine3name",ipsStr)
		if err!=nil {
			log.Println("connect frp server error:",err)
			log.Println("wait 5 sencond...")
			time.Sleep(time.Second*5)
		}
	}
}
