package a

import "fmt"

type TestStruct[T any] struct {
	TestAny        TestInterface
	TestAnyPointer *T
}

type TestInterface interface {
}

func TestFunc[T any](t TestStruct[T], a TestStruct[any]) {

	switch t.TestAny.(type) {
	case string:
		fmt.Printf("hello")
	case map[any]any:
		fmt.Printf("any")
	default:
		fmt.Print("unknown")
	}

	testFun := func() {
		fmt.Printf("hello, %v", a.TestAny.(any))
	}

	testFun()
}
