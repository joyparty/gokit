package gokit

// Must 用于检查错误，如果错误不为空，则抛出异常
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// MustReturn 用于检查错误，如果错误不为空，则抛出异常
//
// Example:
//
//	func Foo() (int, error) {
//		// ...
//	}
//
//	var n int
//	n = MustReturn(Foo())
func MustReturn[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
