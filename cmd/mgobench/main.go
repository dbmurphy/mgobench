package main

import (
	"fmt"

	mgobench "github.com/mgobench"
	flag "github.com/ogier/pflag"
)

// config parser and distributor: remove this function to some other file or package

func main() {
	var configfile string

	flag.StringVarP(&configfile, "config", "c", "", "config file path")

	flag.Parse()

	c, err := mgobench.LoadConfig(configfile)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(c.Thread)
		fmt.Println("======= Mongo Info ========")
		fmt.Println(c.Mongo.ConnectionString)
		fmt.Println(c.Mongo.Database)
		fmt.Println(c.Mongo.Collection)
		fmt.Println("======= Influx Info ========")
		fmt.Println(c.Influxdb.ConnectionString)
		fmt.Println(c.Influxdb.Database)
		fmt.Println(configfile)
		fmt.Println("======= testCase Info ========")

		for serverName, server := range c.Testcases {
			fmt.Println("--------------------*(*(*(*")
			fmt.Printf("Server: %s (%s, %s)\n", serverName, server.Name, server.Duration)
		}
	}

	// Nitin's work
	// wm := NewWorkerManager(3)
	// // fmt.Println(wm.NumWorker())
	// // fmt.Println(wm.IsRunning())

	// ccc := NewCollectionBindFunc("Oorder", "test")
	// mm := MgoManager{
	// 	Session: session,
	// 	CFn:     ccc,
	// }
	// mt := MongoTask{
	// 	SM: mm,
	// }
	// type person struct {
	// 	Name string
	// }
	// var data = make([]interface{}, 0)
	// data = append(data, person{"nitinsanq"})
	// for i := 0; i < 10; i++ {
	// 	ch := InsertTask{
	// 		MongoTask: mt,
	// 		Docs:      data,
	// 		Name:      "Oorder",
	// 	}
	// 	wm.Send(ch)
	// }
	// tr := wm.Result()
	// for d := range tr {
	// 	fmt.Println(d.Count, "time", d.TimeTaken)
}
