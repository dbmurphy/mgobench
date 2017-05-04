package main

import (
	"fmt"

	mgobench "github.com/mgobench"
	flag "github.com/ogier/pflag"
)

func main() {
	fmt.Println("Initiating test")
	var configfile string

	flag.StringVarP(&configfile, "config", "c", "", "config file path")

	flag.Parse()

	c, err := mgobench.LoadConfig(configfile)
	if err != nil {
		panic(err)
	} else {

		fmt.Println("================ Influx tester =========================")
		influxdb := mgobench.NewInfluxClient(&c)
		fmt.Println("1 st of instance ", influxdb)
		fmt.Println("INflux insertiion on")
		fmt.Println()
		influxdb.InsertData("insertTime", "insert_time", 0.00)
		fmt.Println("INflux insertiion off")
	}
}
