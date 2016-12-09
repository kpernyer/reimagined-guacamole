# Testing Plugin from golang 1.8

## Version Local
```
docker run -v "$PWD:/data" golang:1.8beta1-wheezy /data/compile.sh
```

## Local machine and bash into Docker
```
% docker run -it golang:1.8beta1-wheezy bash
```

## In the Docker Bash
```
% go get github.com/kpernyer/reimagined-guacamole/goplug
% cd $GOPATH/src/github.com/kpernyer/reimagined-guacamole/goplug
```

## Build the Plugin
```
# go build -buildmode=plugin -o true_plugin.so myplugin.go`
```

## Build and run the main program
```
# go run main.go
2016/12/09 13:38:45 Two different values give false
2016/12/09 13:38:45 Two same values give true
```
```
# ls
README.md  main.go  myplugin.go  true_plugin.so`
```


