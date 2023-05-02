package main

import (
	"flag"
	"io/ioutil"

	"github.com/msz-kp/go-netdicom/fuzze2e"
)

func main() {
	flag.Parse()
	fuzzData, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	fuzze2e.Fuzz(fuzzData)
}
