rm -r staterecovery
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o staterecovery ./main.go