package mgobench

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"

	mgobench "github.com/mgobench"
	cases "github.com/mgobench/cases"
)

func launchWorker() {

}

func Start(c *mgobench.Config) {

	var session, _ = mgo.Dial(c.Mongo.ConnectionString)
	// newSessCopy := session.Copy()
	// mongo stuff

	// Result Worker Manager
	r := mgobench.NewResultWorker(1, 1*time.Second, c)

	//  Worker Manager
	wm := mgobench.NewWorkerManager(3, r.C)

	// Register each of the Test Cases
	TestCaseRegistry := mgobench.Newregistry()
	TestCaseRegistry.Add("emptyTest", cases.EmptyDocTest)
	TestCaseRegistry.Add("flatT1DocTest", cases.FlatT1DocTest)
	TestCaseRegistry.Add("flatT1InsertTaskTest", cases.FlatT1InsertTaskTest)

	count := 0
	for testcase, test := range c.Testcases {

		// Execute test Cases on diffrent collections
		mt := mgobench.MongoTask{
			SM: mgobench.MgoManager{
				Session: session,
				CFn:     mgobench.NewCollectionBindFunc(c.Mongo.Database, testcase),
			},
			Val: testcase,
		}
		fmt.Printf("Testcases: %s (%s, %s)\n", testcase, test.Name, test.Duration)

		dura, err := time.ParseDuration(test.Duration)
		if err != nil {

			panic(err)
		}
		TestCaseFunc, err := TestCaseRegistry.Get(testcase)
		if err != nil {
			panic(err)
		}

		TestCaseFunc(dura, r, wm, mt)
		count++
		fmt.Printf("-------------------- %s\t - test completed ------------------", testcase)
		fmt.Println("")

	}

	wm.Stop()
	fmt.Println("stopping")
	r.Stop()
	fmt.Println("result stoping")

	// TODO : get test cases from config and map it to function using registry

}
