package main

import (
	flag "github.com/ogier/pflag"
	"fmt"
)


func main() {
	var configfile string
	flag.StringVarP(&configfile, "config", "c", "", "config file path")
	flag.Parse()
	fmt.Println(configfile)
}
