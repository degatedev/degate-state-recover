package main

import (
	"flag"
	"log"
	"staterecovery"
	"staterecovery/conf"
)

func main() {
	log.Println("in main")
	flag.Parse()

	if err := conf.Init(); err != nil {
		panic(err)
	}
	// data from must be 1
	conf.Conf.DataFrom = 1

	staterecovery.StartRecovery(conf.Conf)
}