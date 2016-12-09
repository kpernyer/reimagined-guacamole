# Testing Plugin from golang 1.8

## On local machine
% docker run -it --entrypoint bash golang:1.8beta1-wheezy

1. In the Docker Bash
`% go get github.com/kpernyer/reimagined-guacamole/goplug`
`% cd $GOPATH/src/github.com/kpernyer/reimagined-guacamole/goplug`

2. Build the Plugin
`# go build -buildmode=plugin -o true_plugin.so myplugin.go`

3. Build and run the main program
`# go run main.go
2016/12/09 13:38:45 Two different values give false
2016/12/09 13:38:45 Two same values give true`

`# ls
README.md  main.go  myplugin.go  true_plugin.so`


