package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	a := NewBufferPool(10)
	ccc := NewCollectionBindFunc("ttt", "fdfd")
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
	var interfaceSlice []interface{} = make([]interface{}, 1)
	interfaceSlice[0] = &Person{
		Name: "Nitin",
	}
	aaa := InsertTask{
		MongoTask: mt,
		Docs:      interfaceSlice,
		Name:      "Nitin",
	}
	go func() {
		for i := 0; i < 10; i++ {
			a <- BufferPool{
				T: aaa,
			}
		}
		close(a)
	}()
	for b := range a {
		fmt.Println("&T", b)
		msg := b
		fmt.Println("dffsdf", msg)
		fmt.Println(msg.T.Run())
	}

}
