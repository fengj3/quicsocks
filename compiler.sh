#!/bin/bash
compiler() {
  if [ "$1" == "windows" -o "$1" == "darwin" ];then
    suffix=".exe";
  else
    suffix="";
  fi
  CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -o bin/quicsocks-server-${1}-${2}${suffix} server/server.go
  CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -o bin/quicsocks-client-${1}-${2}${suffix} client/client.go
  CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -o bin/quicsocks-client-plus-${1}-${2}${suffix} client-plus/client-plus.go
}
compiler linux amd64
compiler linux 386
compiler linux arm64
compiler linux arm
compiler windows amd64
compiler windows 386
compiler darwin amd64
#compiler darwin 386 #unsupported GOOS/GOARCH pair darwin/386