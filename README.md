# rpcx-case

## 初始化
- export GO111MODULE=on
- export GOPROXY=https://mirrors.aliyun.com/goproxy/
- go mod init rpcx-case [用于项目 go mod 初始化]

## 安装rpcx 
-  go get -u -v -tags "reuseport quic kcp zookeeper etcd consul ping" github.com/smallnest/rpcx/...

