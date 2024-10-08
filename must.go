package gokit

import (
	"fmt"
)

// Must 用于检查错误，如果错误不为空，则抛出异常
func Must(errors ...error) {
	for _, err := range errors {
		if err != nil {
			panic(err)
		}
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

// MustReturn2 用于检查两个有效返回值时的错误，如果错误不为空，则抛出异常
//
// Example:
//
//	func Foo() (int, string, error) {
//		// ...
//	}
//
//	var n int
//	var s string
//	n, s = MustReturn2(Foo())
func MustReturn2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return t1, t2
}

// MustTo 类型转换失败则panic
//
// Deprecated: Use MustBe() instead.
func MustTo[T any](v any) T {
	return MustBe[T](v)
}

// MustBe 类型断言，失败则panic
func MustBe[T any](v any) T {
	if x, ok := v.(T); ok {
		return x
	}

	var x T
	panic(fmt.Errorf("%T is not %T", v, x))
}
