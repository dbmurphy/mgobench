package main

import (
	"fmt"
)

func main() {
	wm := NewWorkerManager(3)
	// fmt.Println(wm.NumWorker())
	// fmt.Println(wm.IsRunning())

	ccc := NewCollectionBindFunc("Oorder", "test")
	mm := MgoManager{
		Session: session,
		CFn:     ccc,
	}
	mt := MongoTask{
		SM: mm,
	}
	type person struct {
		Name string
	}
	var data = make([]interface{}, 0)
	data = append(data, person{"nitinsanq"})
	for i := 0; i < 10; i++ {
		ch := InsertTask{
			MongoTask: mt,
			Docs:      data,
			Name:      "Oorder",
		}
		wm.Send(ch)
	}
	tr := wm.Result()
	for d := range tr {
		fmt.Println(d.Count, "time", d.TimeTaken)
	}
}
