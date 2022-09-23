#! /bin/bash

rm -rf build

[ ! -d build ] && mkdir build
[ ! -d build/darwin/amd64 ] && mkdir -p build/darwin/amd64
[ ! -d build/linux/amd64 ] && mkdir -p build/linux/amd64
[ ! -d build/windows/amd64 ] && mkdir -p build/windows/amd64

TARGET_OS_LIST="windows linux darwin"
TARGET_ARCH_LIST="amd64"

for GOOS in $TARGET_OS_LIST
do
  for GOARCH in $TARGET_ARCH_LIST
  do
    if [ "$GOOS" == "windows" ]
    then
      BINNAME=tdg.exe
    else
      BINNAME=tdg
    fi
    echo build ${GOOS}/${GOARCH}/${BINNAME}
    GOOS=$GOOS GOARCH=$GOARCH go build -o build/${GOOS}/${GOARCH}/${BINNAME} cmd/tdg/main.go
  done
done
