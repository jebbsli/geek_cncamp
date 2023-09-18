export CGO_ENABLED=0
export GO111MODULE=on
export GOOS=linux
export GOARCH=amd64

GO_LDFLAGS += -extldflags '-static'

build:
	go build --trimpath -tags=jsoniter -ldflags="${GO_LDFLAGS}"  -o ./bin/server ./server.go

image.%:
	make build
	docker build --platform linux/amd64 -f Dockerfile -t jebbsli/geek_cncamp:${*} .

push.%:
	make image.${*}
	docker push jebbsli/geek_cncamp:${*}