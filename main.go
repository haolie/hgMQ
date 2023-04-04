package main

import (
	"github.com/haolie/goUtil/logUtil"
	"lyh/hgMQ/config"
	"lyh/hgMQ/sys"
)

func main() {
	config.LoadConfig()
	logUtil.InitLog()

	ctx, success := sys.Start()
	if success == false {
		<-ctx.Done()
	}
}
