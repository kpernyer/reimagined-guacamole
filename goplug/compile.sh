#! /bin/sh

cd /data
go build -buildmode=plugin -o true_plugin.so myplugin.go
go run main.go
