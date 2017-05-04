package main

import (
	"fmt"
	"sync"
	"time"

	mgobench "github.com/mgobench"
	flag "github.com/ogier/pflag"
)

func main() {
	var wg sync.WaitGroup
	// Config test pass
	fmt.Println("Initiating test")
	var configfile string

	flag.StringVarP(&configfile, "config", "c", "", "config file path")

	flag.Parse()

	c, err := mgobench.LoadConfig(configfile)
	if err != nil {
		panic(err)
	} else {

		// session, _ := mgo.Dial(c.Mongo.ConnectionString)
		// Result Worker Manager
		r := mgobench.NewResultWorker(1, 1*time.Second, &c)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				r.C <- mgobench.TaskResult{
					Count:     i,
					TimeTaken: time.Since(time.Now()),
				}
				fmt.Println(i)
				time.Sleep(100 * time.Millisecond)
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 15; i++ {
				r.C <- mgobench.TaskResult{
					Count:     i,
					TimeTaken: time.Since(time.Now()),
				}
				fmt.Println(i)
				time.Sleep(100 * time.Millisecond)
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 25; i++ {
				r.C <- mgobench.TaskResult{
					Count:     i,
					TimeTaken: time.Since(time.Now()),
				}
				fmt.Println(i)
				time.Sleep(100 * time.Millisecond)
			}
		}()
		wg.Wait()
		r.Stop()

	}
	// 	fmt.Println("================ Influx tester =========================")
	// 	influxdb := mgobench.NewInfluxClient(&c)
	// 	fmt.Println("1 st of instance ", influxdb)
	// 	fmt.Println("INflux insertiion on")
	// 	fmt.Println()
	// 	influxdb.InsertData("insertTime", "insert_time", 0.00)
	// 	fmt.Println("INflux insertiion off")
	// }
}
