package SFStringsUtil

import (
	// "math"
	"strconv"
	"testing"
)

/**
 *  test string to int
 */
func TestToInt(t *testing.T) {

	// math.MaxInt8 = 127
	// math.MinInt8 = -128
	// math.MaxInt16 = 32767
	// math.MinInt16 = -32768
	// math.MaxInt32 = 2147483647
	// math.MinInt32 = -2147483648
	// math.MaxInt64 =  9223372036854775807
	// math.MinInt64 = -9223372036854775808

	tint8 := ToInt8("127")
	if tint8 != 127 {
		t.Fatal("Fatal: ToInt8(\"127\")", tint8)
	}
	tint8 = ToInt8("128")
	if tint8 != 127 {
		t.Fatal("Fatal: ToInt8(\"128\")", tint8)
	}
	tint8 = ToInt8("-128")
	if tint8 != -128 {
		t.Fatal("Fatal: ToInt8(\"-128\")", tint8)
	}
	tint8 = ToInt8("-129")
	if tint8 != -128 {
		t.Fatal("Fatal: ToInt8(\"-129\")", tint8)
	}
	tint8 = ToInt8("0")
	if tint8 != 0 {
		t.Fatal("Fatal: ToInt8(\"0\")", tint8)
	}

	tint16 := ToInt16("32767")
	if tint16 != 32767 {
		t.Fatal("Fatal: ToInt16(\"32767\")", tint16)
	}
	tint16 = ToInt16("32768")
	if tint16 != 32767 {
		t.Fatal("Fatal: ToInt16(\"32768\")", tint16)
	}
	tint16 = ToInt16("-32768")
	if tint16 != -32768 {
		t.Fatal("Fatal: ToInt16(\"-32768\")", tint16)
	}
	tint16 = ToInt16("-32769")
	if tint16 != -32768 {
		t.Fatal("Fatal: ToInt16(\"-32769\")", tint16)
	}
	tint16 = ToInt16("0")
	if tint16 != 0 {
		t.Fatal("Fatal: ToInt16(\"0\")", tint16)
	}

	tint32 := ToInt32("2147483647")
	if tint32 != 2147483647 {
		t.Fatal("Fatal: ToInt32(\"2147483647\")", tint32)
	}
	tint32 = ToInt32("2147483648")
	if tint32 != 2147483647 {
		t.Fatal("Fatal: ToInt32(\"2147483648\")", tint32)
	}
	tint32 = ToInt32("-2147483648")
	if tint32 != -2147483648 {
		t.Fatal("Fatal: ToInt32(\"-2147483648\")", tint32)
	}
	tint32 = ToInt32("-2147483649")
	if tint32 != -2147483648 {
		t.Fatal("Fatal: ToInt32(\"-2147483649\")", tint32)
	}
	tint32 = ToInt32("0")
	if tint32 != 0 {
		t.Fatal("Fatal: ToInt32(\"0\")", tint32)
	}

	tint64 := ToInt64("9223372036854775807")
	if tint64 != 9223372036854775807 {
		t.Fatal("Fatal: ToInt32(\"9223372036854775807\")", tint64)
	}
	tint64 = ToInt64("9223372036854775809")
	if tint64 != 9223372036854775807 {
		t.Fatal("Fatal: ToInt32(\"9223372036854775809\")", tint64)
	}
	tint64 = ToInt64("-9223372036854775808")
	if tint64 != -9223372036854775808 {
		t.Fatal("Fatal: ToInt32(\"-9223372036854775808\")", tint64)
	}
	tint64 = ToInt64("-9423372036854775809")
	if tint64 != -9223372036854775808 {
		t.Fatal("Fatal: ToInt32(\"-9423372036854775809\")", tint64)
	}
	tint64 = ToInt64("0")
	if tint64 != 0 {
		t.Fatal("Fatal: ToInt32(\"0\")", tint64)
	}

	tint := ToInt("9223372036854775807")
	if tint != 9223372036854775807 {
		t.Fatal("Fatal: ToInt(\"9223372036854775807\")", tint)
	}

	tint = ToInt("8")
	if tint != 8 {
		t.Fatal("Fatal: ToInt(\"8\")", tint)
	}

	tint = ToInt("+8")
	if tint != 8 {
		t.Fatal("Fatal: ToInt(\"+8\")", tint)
	}

	tint = ToInt("-8")
	if tint != -8 {
		t.Fatal("Fatal: ToInt(\"-8\")", tint)
	}

	tint = ToInt("-8yrmp+")
	if tint != -8 {
		t.Fatal("Fatal: ToInt(\"-8yrmp+\")", tint)
	}

	tint = ToInt("y-8yrmp+")
	if tint != 0 {
		t.Fatal("Fatal: ToInt(\"y-8yrmp+\")", tint)
	}

	tint = ToInt("  -8")
	if tint != -8 {
		t.Fatal("Fatal: ToInt(\"  -8\")", tint)
	}

	tint = ToInt("  888-")
	if tint != 888 {
		t.Fatal("Fatal: ToInt(\"  888-\")", tint)
	}
}

/**
 *   ToInt 与 strconv.Atoi 比较
 */
func Benchmark_int_ToInt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		strint := strconv.Itoa(i)
		tint := ToInt(strint)
		if tint != i {
			b.Fatalf("Fail: ToInt(%#v)", strint)
		}
	}
}

/**
 *  ToInt 与 strconv.Atoi 比较
 */
func Benchmark_int_Atoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strint := strconv.Itoa(i)
		tint, _ := strconv.Atoi(strint)
		if tint != i {
			b.Fatalf("Fail: strconv.Atoi(%#v)", strint)
		}
	}
}
