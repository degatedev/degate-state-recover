if [ -f staterecovery ]
then
    rm -f staterecovery
fi
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o staterecovery ./main.go