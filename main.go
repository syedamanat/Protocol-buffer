package main

import (
	"fmt"

	simplepb "github.com/syedamanat/protobuf-example-go/src/simple"
)

func main() {
	fmt.Println("Hello Worlld")
	sm := simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "Syed",
		SampleList: []int32{1, 3, 5, 7},
	}

	fmt.Println(sm)

	sm.Name = "New Name"
	fmt.Println(sm)

	fmt.Println("The ID of person is ", sm.GetId())

}
