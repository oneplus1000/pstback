package main

import (
	"flag"
	"log"

	"github.com/oneplus1000/pstback/pstback"
)

var cfgpath = flag.String("config", "", "config file")

func main() {
	flag.Parse()
	err := pstback.BackUp(*cfgpath)
	if err != nil {
		echoErr(err)
	}
}

func echoErr(err error) {
	log.Panicf("%+v", err)
}
