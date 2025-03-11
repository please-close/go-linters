package a

import "fmt"

type TestStruct[T any] struct {
	TestAny        TestInterface
	TestAnyPointer *T
}

type TestInterface interface {
}

//nolint:detectany
func TestFunc[T any](t TestStruct[T], a TestStruct[any]) {
	// do nothing
	var anyMap map[any]any

	anyMap["123"] = 123
	anyMap["13"] = "123"

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
