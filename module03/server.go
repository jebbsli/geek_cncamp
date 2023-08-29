package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

/*
 * 编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

 * 1. 接收客户端 request，并将 request 中带的 header 写入 response header
 * 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
 * 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
 * 4. 当访问 localhost/healthz 时，应返回 200
 */

func init() {
	err := os.Setenv("VERSION", time.Now().Format("20060102150406"))
	if err != nil {
		log.Fatalf("set env error: %s", err.Error())
	}
}

func HealthZ(writer http.ResponseWriter, request *http.Request) {
	// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
	for k, vs := range request.Header {
		for _, v := range vs {
			writer.Header().Add(k, v)
		}
	}

	// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	writer.Header().Set("VERSION", os.Getenv("VERSION"))

	// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Printf("clientIp: %s, response code: %d", request.Host, http.StatusOK)

	// 4. 当访问 localhost/healthz 时，应返回 200
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("the server is healthy"))
}

func main() {
	http.HandleFunc("/healthz", HealthZ)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("start server error: %s", err.Error())
	}
}
