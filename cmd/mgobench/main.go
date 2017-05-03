package main

import (
	"runtime"

	mgobench "github.com/mgobench"
	launcher "github.com/mgobench/launcher"
	flag "github.com/ogier/pflag"
)

// config parser and distributor: remove this function to some other file or package

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var configfile string

	flag.StringVarP(&configfile, "config", "c", "", "config file path")

	flag.Parse()

	c, err := mgobench.LoadConfig(configfile)
	if err != nil {
		panic(err)
	} else {
		launcher.Start(&c)
	}
}
