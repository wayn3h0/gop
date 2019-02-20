package testing

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const (
	// EqualFormat is the string format for equal test.
	EqualFormat = "\n%s:%d: should be == \n   \thave: (%T) %+v\n\twant: (%T) %+v"
	// EqualLFormat is the string format for equal loop test.
	EqualLFormat = "\n%s:%d: should be == \n%d.\thave: (%T) %+v\n\twant: (%T) %+v"
	// NotEqualFormat is the string format for not equal test.
	NotEqualFormat = "\n%s:%d: should be != \n   \thave: (%T) %+v\n\twant not: (%T) %+v"
	// NotEqualLFormat is the string format for not euqal loop test.
	NotEqualLFormat = "\n%s:%d: should be != \n%d.\thave: (%T) %+v\n\twant not: (%T) %+v"
)

// caller returns file and line two stack frames.
func caller() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}

	return file, line
}

// ExpectEqual calls t.Error and prints a nice comparison message when have != want.
func ExpectEqual(tb testing.TB, have, want interface{}) {
	if !reflect.DeepEqual(have, want) {
		file, line := caller()
		tb.Errorf(EqualFormat, file, line, have, have, want, want)
	}
}

// ExpectNotEqual calls t.Error and prints a nice comparison message when have == want.
func ExpectNotEqual(tb testing.TB, have, wantnot interface{}) {
	if reflect.DeepEqual(have, wantnot) {
		file, line := caller()
		tb.Fatalf(NotEqualFormat, file, line, have, have, wantnot, wantnot)
	}
}

// AssertEqual calls t.Fatal to abort the test immediately and prints a nice comparison message when have != want.
func AssertEqual(tb testing.TB, have, want interface{}) {
	if !reflect.DeepEqual(have, want) {
		file, line := caller()
		tb.Fatalf(EqualFormat, file, line, have, have, want, want)
	}
}

// AssertNotEqual calls t.Fatal to abort the test immediately and prints a nice comparison message when have == wantnot.
func AssertNotEqual(tb testing.TB, have, wantnot interface{}) {
	if reflect.DeepEqual(have, wantnot) {
		file, line := caller()
		tb.Fatalf(NotEqualFormat, file, line, have, have, wantnot, wantnot)
	}
}

// ExpectEqualL similar to ExpectEqual, it's used for loop (table-based) tests.
func ExpectEqualL(tb testing.TB, have, want interface{}, index int) {
	if !reflect.DeepEqual(have, want) {
		file, line := caller()
		tb.Errorf(EqualLFormat, file, line, index, have, have, want, want)
	}
}

// ExpectNotEqualL similar to ExpectNotEqual, it's used for loop (table-based) tests.
func ExpectNotEqualL(tb testing.TB, have, wantnot interface{}, index int) {
	if reflect.DeepEqual(have, wantnot) {
		file, line := caller()
		tb.Fatalf(NotEqualLFormat, file, line, index, have, have, wantnot, wantnot)
	}
}

// AssertEqualL similar to AssertEqual, it's used for loop (table-based) tests.
func AssertEqualL(tb testing.TB, have, want interface{}, index int) {
	if !reflect.DeepEqual(have, want) {
		file, line := caller()
		tb.Fatalf(EqualLFormat, file, line, index, have, have, want, want)
	}
}

// AssertNotEqualL similar to AssertEqual, it's used for loop (table-based) tests.
func AssertNotEqualL(tb testing.TB, have, wantnot interface{}, index int) {
	if reflect.DeepEqual(have, wantnot) {
		file, line := caller()
		tb.Fatalf(NotEqualLFormat, file, line, index, have, have, wantnot, wantnot)
	}
}
