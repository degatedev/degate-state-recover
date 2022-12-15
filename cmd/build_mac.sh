if [ -f staterecovery ]
then
    rm -f staterecovery
fi
go build -o staterecovery ./main.go