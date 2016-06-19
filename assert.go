package assert

// Testing helpers for doozer.

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/kr/pretty"
)

func assert(t *testing.T, result bool, f func(), cd int) {
	if !result {
		_, file, line, _ := runtime.Caller(cd + 1)
		t.Errorf("%s:%d", file, line)
		f()
		t.FailNow()
	}
}

func equal(t *testing.T, exp, got interface{}, cd int, args ...interface{}) {
	fn := func() {
		for _, desc := range pretty.Diff(exp, got) {
			t.Error("!", desc)
		}
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	result := reflect.DeepEqual(exp, got)
	assert(t, result, fn, cd+1)
}

func tt(t *testing.T, result bool, cd int, args ...interface{}) {
	fn := func() {
		t.Errorf("!  Failure")
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	assert(t, result, fn, cd+1)
}

func True(t *testing.T, result bool, args ...interface{}) {
	tt(t, result, 1, args...)
}

func Truef(t *testing.T, result bool, format string, args ...interface{}) {
	tt(t, result, 1, fmt.Sprintf(format, args...))
}

//func True(t *testing.T, boolValue bool, logs ...interface{}) {
//	trueAssert(t, false, boolValue, logs...)
//}

func MustTrue(t *testing.T, boolValue bool, logs ...interface{}) {
	trueAssert(t, true, boolValue, logs...)
}

func trueAssert(t *testing.T, fatal bool, value bool, logs ...interface{}) {
	if !value {
		logCaller(t)
		if len(logs) > 0 {
			t.Log(logs...)
		} else {
			t.Logf("value is not true")
		}
		failIt(t, fatal)
	}
}

func Panic(t *testing.T, err interface{}, fn func()) {
	defer func() {
		equal(t, err, recover(), 3)
	}()
	fn()
}

func Nil(t *testing.T, value interface{}, logs ...interface{}) {
	nilAssert(t, false, true, value, logs...)
}

func MustNil(t *testing.T, value interface{}, logs ...interface{}) {
	nilAssert(t, true, true, value, logs...)
}

func NotNil(t *testing.T, value interface{}, logs ...interface{}) {
	nilAssert(t, false, false, value, logs...)
}

func MustNotNil(t *testing.T, value interface{}, logs ...interface{}) {
	nilAssert(t, true, false, value, logs...)
}

func nilAssert(t *testing.T, fatal bool, isNil bool, value interface{}, logs ...interface{}) {
	if isNil != (value == nil || reflect.ValueOf(value).IsNil()) {
		logCaller(t)
		if len(logs) > 0 {
			t.Log(logs...)
		} else {
			if isNil {
				t.Log("value is not nil:", value)
			} else {
				t.Log("value is nil")
			}
		}
		failIt(t, fatal)
	}
}

func Equal(t *testing.T, exp, got interface{}, args ...interface{}) {
	equal(t, exp, got, 1, args...)
}

func Equalf(t *testing.T, exp, got interface{}, format string, args ...interface{}) {
	equal(t, exp, got, 1, fmt.Sprintf(format, args...))
}

func NotEqual(t *testing.T, exp, got interface{}, args ...interface{}) {
	fn := func() {
		t.Errorf("!  Unexpected: <%#v>", exp)
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	result := !reflect.DeepEqual(exp, got)
	assert(t, result, fn, 1)
}

func Equals(t *testing.T, expected, actual interface{}, logs ...interface{}) {
	equalAssert(t, false, true, expected, actual, logs...)
}
func MustEqual(t *testing.T, expected, actual interface{}, logs ...interface{}) {
	equalAssert(t, true, true, expected, actual, logs...)
}

func NotEquals(t *testing.T, expected, actual interface{}, logs ...interface{}) {
	equalAssert(t, false, false, expected, actual, logs...)
}

func MustNotEqual(t *testing.T, expected, actual interface{}, logs ...interface{}) {
	equalAssert(t, true, false, expected, actual, logs...)
}

func EqualSprint(t *testing.T, expected, actual interface{}, logs ...interface{}) {
	equalAssert(t, false, true, fmt.Sprint(expected), fmt.Sprint(actual), logs...)
}

func MustEqualSprint(t *testing.T, expected, actual interface{}, logs ...interface{}) {
	equalAssert(t, true, true, fmt.Sprint(expected), fmt.Sprint(actual), logs...)
}

func NotEqualSprint(t *testing.T, expected, actual interface{}, logs ...interface{}) {
	equalAssert(t, false, false, fmt.Sprint(expected), fmt.Sprint(actual), logs...)
}

func MustNotEqualSprint(t *testing.T, expected, actual interface{}, logs ...interface{}) {
	equalAssert(t, true, false, fmt.Sprint(expected), fmt.Sprint(actual), logs...)
}

func equalAssert(t *testing.T, fatal bool, isEqual bool, expected, actual interface{}, logs ...interface{}) {
	expected = normalizeValue(expected)
	actual = normalizeValue(actual)
	if isEqual != (reflect.DeepEqual(expected, actual)) {
		logCaller(t)
		if len(logs) > 0 {
			t.Log(logs...)
		} else {
			if isEqual {
				t.Log("Values not equal")
			} else {
				t.Log("Values equal")
			}
		}
		t.Log("Expected: ", expected)
		t.Log("Actual: ", actual)
		failIt(t, fatal)
	}
}

func Zero(t *testing.T, value interface{}, logs ...interface{}) {
	zeroAssert(t, false, true, false, value, logs...)
}
func MustZero(t *testing.T, value interface{}, logs ...interface{}) {
	zeroAssert(t, true, true, false, value, logs...)
}
func NotZero(t *testing.T, value interface{}, logs ...interface{}) {
	zeroAssert(t, false, false, false, value, logs...)
}
func MustNotZero(t *testing.T, value interface{}, logs ...interface{}) {
	zeroAssert(t, true, false, false, value, logs...)
}

func ZeroLen(t *testing.T, value interface{}, logs ...interface{}) {
	zeroAssert(t, false, true, true, value, logs...)
}

func MustZeroLen(t *testing.T, value interface{}, logs ...interface{}) {
	zeroAssert(t, true, true, true, value, logs...)
}

func PositiveLen(t *testing.T, value interface{}, logs ...interface{}) {
	zeroAssert(t, false, false, true, value, logs...)
}

func MustPositiveLen(t *testing.T, value interface{}, logs ...interface{}) {
	zeroAssert(t, true, false, true, value, logs...)
}

func zeroAssert(t *testing.T, fatal bool, isZero bool, length bool, value interface{}, logs ...interface{}) {
	var name string
	var integerValue int
	value = normalizeValue(value)
	v := reflect.Indirect(reflect.ValueOf(value))
	if length {
		name = "Length"
		integerValue = v.Len()
	} else {
		name = "Value"
		integerValue = int(v.Int())
	}
	if isZero != (integerValue == 0) {
		logCaller(t)
		if len(logs) > 0 {
			t.Log(logs...)
		} else {
			if isZero {
				t.Log(name, "is not zero:", value)
			} else {
				t.Log(name, "is zero.")
			}
		}
		failIt(t, fatal)
	}
}

func OneLen(t *testing.T, value interface{}, logs ...interface{}) {
	oneLenAssert(t, false, value, logs...)
}

func MustOneLen(t *testing.T, value interface{}, logs ...interface{}) {
	oneLenAssert(t, true, value, logs...)
}

func oneLenAssert(t *testing.T, fatal bool, value interface{}, logs ...interface{}) {
	v := reflect.Indirect(reflect.ValueOf(value))
	if v.Len() != 1 {
		logCaller(t)
		if len(logs) > 0 {
			t.Log(logs...)
		} else {
			t.Log("Length is not one:", v.Len())
		}
		failIt(t, fatal)
	}
}

func logCaller(t *testing.T) {
	_, file, line, _ := runtime.Caller(3)
	t.Logf("Caller: %v:%d", file, line)
}

func failIt(t *testing.T, fatal bool) {
	if fatal {
		t.FailNow()
	} else {
		t.Fail()
	}
}

func normalizeValue(value interface{}) interface{} {
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(val.Uint())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int()
	case reflect.Float32, reflect.Float64:
		return val.Float()
	case reflect.Complex64, reflect.Complex128:
		return val.Complex()
	case reflect.String:
		return val.String()
	case reflect.Bool:
		return val.Bool()
	case reflect.Slice:
		if val.Type().Elem().Kind() == reflect.Uint8 {
			return val.Bytes()
		}
	}
	return value
}
