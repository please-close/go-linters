package testdata

type TestStruct[T any] struct {
	TestAny        any
	TestAnyPointer *any
}

func TestFunc[T any](t T, a any) {
	// do nothing
}
