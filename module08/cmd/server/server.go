package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
	"time"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("current time: ", fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05")))
	})
	serverAddr, err := g.Cfg().Get(context.TODO(), "server.address")
	if err != nil {
		log.Fatalf("get server address failed, error: %s", err.Error())
	}
	s.SetAddr(serverAddr.String())
	s.Run()
}
