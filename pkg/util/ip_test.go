package util

import (
	"log"
	"testing"
)

func TestExternalIP(t *testing.T) {
	ips := ExternalIP()

	if ips!=nil && len(ips)>0{

		log.Println(ips)
	}
}