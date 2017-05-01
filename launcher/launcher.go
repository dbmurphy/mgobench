package mgobench

import (
	"fmt"
	"time"

	mgobench "github.com/mgobench"
	cases "github.com/mgobench/cases"
	trw "github.com/mgobench/taskResultWorker"
	mgo "gopkg.in/mgo.v2"
)

func Start(c *mgobench.Config) {

	fmt.Println(c.Thread)
	fmt.Println("======= Mongo Info ========")
	fmt.Println(c.Mongo.ConnectionString)
	fmt.Println(c.Mongo.Database)
	fmt.Println(c.Mongo.Collection)
	fmt.Println("======= Influx Info ========")
	fmt.Println(c.Influxdb.ConnectionString)
	fmt.Println(c.Influxdb.Database)
	fmt.Println("======= testCase Info ========")
	for testcase, test := range c.Testcases {

		fmt.Printf("Server: %s (%s, %s)\n", testcase, test.Name, test.Duration)
	}

	// for testcase, case := range c.Testcases {
	// 	fmt.Printf("Server: %s (%s, %s)\n", testcase, case.Name, case.Duration)
	// 	// getRegisttry(case.Name)
	// 	// for time.After(case.Duration){

	// 	// 	fdsfsafsd...CollectionBindFunc
	// 	// 	time.Sleep(15 * time.Minuet)
	// 	// }
	// }

	// mongo stuff

	var session, _ = mgo.Dial(c.Mongo.ConnectionString) // handle error replace _ with handler for err
	ccc := mgobench.NewCollectionBindFunc(c.Mongo.Database, c.Mongo.Collection)
	mm := mgobench.MgoManager{
		Session: session,
		CFn:     ccc,
	}
	mt := mgobench.MongoTask{
		SM: mm,
	}

	// Result Worker Manager
	r := trw.NewResultWorker(5, 2*time.Second, c)

	//  Worker Manager
	wm := mgobench.NewWorkerManager(3, r.C)

	// TODO : get test cases from config and map it to function using registry
	type person struct {
		Name string
	}
	var data = make([]interface{}, 0)
	data = append(data, cases.EmptyTest())
	ch := mgobench.InsertTask{
		MongoTask: mt,
		Docs:      data,
		Name:      "Oorder",
	}

Loop:
	for {
		select {

		case <-time.After(10 * time.Minute):
			// send to influxdb
			break Loop
		default:
			wm.Send(ch)
		}
	}
	wm.Stop()
	r.Stop()

	// TODO: optimize result channel
	// tr := wm.Result()  // No need to have this

	// for d := range tr {
	// 	r.C <- d
	// 	//defer d.Session.Close()
	// 	fmt.Println(d.Count, "time", d.TimeTaken)
	// }
	// wm.Stop()
	// r.Stop()
}
