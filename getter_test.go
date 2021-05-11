package mconfig

import (
	"fmt"
	"strings"
	"testing"

	"mlib.com/mlog"
)

func TestGetString(t *testing.T) {
	mlog.SetLogDir("./logs")
	SetConfigFile("./config-getstring-test.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	}
	for i := 1; i < 9; i++ {
		key := fmt.Sprintf("A%d", i)
		_, err := GetString(key)
		if err != nil {
			t.Errorf("GetString key=%s failed", key)
		}
	}
}

func TestGetBool(t *testing.T) {
	mlog.SetLogDir("./logs")
	SetConfigFile("./config-getstring-test.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	}
	for i := 1; i < 9; i++ {
		key := fmt.Sprintf("A%d", i)
		_, err := GetBool(key)
		if err != nil {
			t.Errorf("GetString key=%s failed", key)
		}
	}

	WriteConfig()
}

func TestGetInt(t *testing.T) {
	testKeys := []string{
		"A-int16-1",
		"A-int16-2",
		"A-int16-3",
		"A-int16-4",
		"A-int16-5",
		"A-int16-6",
		"A-int16-7",
		"A-int16-8",
		"A-int32-1",
		"A-int32-2",
		"A-int32-3",
		"A-int32-4",
		"A-int32-5",
		"A-int32-6",
		"A-int32-7",
		"A-int32-8",
		"A-int64-1",
		"A-int64-2",
		"A-int64-3",
		"A-int64-4",
		"A-int64-5",
		"A-int64-6",
		"A-int64-7",
		"A-int64-8",
		"A-int8-1",
		"A-int8-2",
		"A-int8-3",
		"A-int8-4",
		"A-int8-5",
		"A-int8-6",
		"A-int8-7",
		"A-int8-8",
		"A-uint16-1",
		"A-uint16-2",
		"A-uint16-3",
		"A-uint16-4",
		"A-uint32-1",
		"A-uint32-2",
		"A-uint32-3",
		"A-uint32-4",
		"A-uint64-1",
		"A-uint64-2",
		"A-uint64-3",
		"A-uint64-4",
		"A-uint8-1",
		"A-uint8-2",
		"A-uint8-3",
		"A-uint8-4",
	}
	testExpect := []interface{}{
		32767,
		32767,
		32768,
		32768,
		-32768,
		-32768,
		-32769,
		-32769,
		2147483647,
		2147483647,
		0,
		0,
		0,
		-2147483648,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		127,
		127,
		128,
		128,
		-128,
		-128,
		-129,
		-129,
		65535,
		65535,
		65536,
		65536,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		255,
		255,
		256,
		256,
	}
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-getint.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	}
	for i := 1; i < len(testKeys); i++ {
		key := fmt.Sprintf("A%d", i)
		v, err := GetInt(testKeys[i])
		if err != nil {
			t.Errorf("GetInt key=%s failed:%v", key, err)
		} else {
			if v != testExpect[i] {
				t.Errorf("testKey=%v get=%v testExpect=%v failed", testKeys[i], v, testExpect[i])
			}
		}
	}
	_, err = GetInt("A-test-string-invalid")
	if err == nil {
		t.Errorf("GetInt key=A-test-string-invalid failed")
	}

	v1, err := GetInt("A-test-bool-true")
	if err != nil {
		t.Errorf("GetInt key=A-test-bool-true failed:%v", err)
	} else {
		if v1 != 1 {
			t.Errorf("GetInt key=A-test-bool-true failed:expect 1,but get v1=%v", v1)
		}
	}
	v2, err := GetInt("A-test-bool-false")
	if err != nil {
		t.Errorf("GetInt key=A-test-bool-false failed:%v", err)
	} else {
		if v2 != 0 {
			t.Errorf("GetInt key=A-test-bool-false failed:expect 0,but get v1=%v", v1)
		}
	}
}

func TestGetUint(t *testing.T) {
	testKeys := []string{
		"A-int16-1",
		"A-int16-2",
		"A-int16-3",
		"A-int16-4",
		"A-int16-5",
		"A-int16-6",
		"A-int16-7",
		"A-int16-8",
		"A-int32-1",
		"A-int32-2",
		"A-int32-3",
		"A-int32-4",
		"A-int32-5",
		"A-int32-6",
		"A-int32-7",
		"A-int32-8",
		"A-int64-1",
		"A-int64-2",
		"A-int64-3",
		"A-int64-4",
		"A-int64-5",
		"A-int64-6",
		"A-int64-7",
		"A-int64-8",
		"A-int8-1",
		"A-int8-2",
		"A-int8-3",
		"A-int8-4",
		"A-int8-5",
		"A-int8-6",
		"A-int8-7",
		"A-int8-8",
		"A-uint16-1",
		"A-uint16-2",
		"A-uint16-3",
		"A-uint16-4",
		"A-uint32-1",
		"A-uint32-2",
		"A-uint32-3",
		"A-uint32-4",
		"A-uint64-1",
		"A-uint64-2",
		"A-uint64-3",
		"A-uint64-4",
		"A-uint8-1",
		"A-uint8-2",
		"A-uint8-3",
		"A-uint8-4",
	}
	testExpect := []uint{
		32767,
		32767,
		32768,
		32768,
		0,
		0,
		0,
		0,
		2147483647,
		2147483647,
		2147483648,
		2147483648,
		0,
		0,
		0,
		0,
		9223372036854775807,
		9223372036854775807,
		9223372036854775807,
		9223372036854775807,
		0,
		0,
		0,
		0,
		127,
		127,
		128,
		128,
		0,
		0,
		0,
		0,
		65535,
		65535,
		65536,
		65536,
		4294967295,
		4294967295,
		4294967296,
		4294967296,
		9223372036854775807,
		9223372036854775807,
		9223372036854775807,
		9223372036854775807,
		255,
		255,
		256,
		256,
	}
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-getint.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	}
	for i := 1; i < len(testKeys); i++ {
		v, err := GetUint(testKeys[i])
		if err != nil {
			t.Errorf("GetUint key=%s failed:%v", testKeys[i], err)
		} else {
			if v != testExpect[i] {
				t.Errorf("testKey=%v get=%v testExpect=%v failed", testKeys[i], v, testExpect[i])
			}
		}
	}
	_, err = GetUint("A-test-string-invalid")
	if err == nil {
		t.Errorf("GetUint key=A-test-string-invalid failed")
	}

	v1, err := GetUint("A-test-bool-true")
	if err != nil {
		t.Errorf("GetUint key=A-test-bool-true failed:%v", err)
	} else {
		if v1 != 1 {
			t.Errorf("GetUint key=A-test-bool-true failed:expect 1,but get v1=%v", v1)
		}
	}
	v2, err := GetUint("A-test-bool-false")
	if err != nil {
		t.Errorf("GetUint key=A-test-bool-false failed:%v", err)
	} else {
		if v2 != 0 {
			t.Errorf("GetUint key=A-test-bool-false failed:expect 0,but get v1=%v", v1)
		}
	}
}

func TestGetInt32(t *testing.T) {
	testMap := map[string]int32{
		"A-int16-1":  32767,
		"A-int16-2":  32767,
		"A-int16-3":  32768,
		"A-int16-4":  32768,
		"A-int16-5":  -32768,
		"A-int16-6":  -32768,
		"A-int16-7":  -32769,
		"A-int16-8":  -32769,
		"A-int32-1":  2147483647,
		"A-int32-2":  2147483647,
		"A-int32-3":  -2147483648,
		"A-int32-4":  -2147483648,
		"A-int32-5":  -2147483648,
		"A-int32-6":  -2147483648,
		"A-int32-7":  -2147483648,
		"A-int32-8":  2147483647,
		"A-int64-1":  -2147483648,
		"A-int64-2":  -1,
		"A-int64-3":  -2147483648,
		"A-int64-4":  -1,
		"A-int64-5":  -2147483648,
		"A-int64-6":  0,
		"A-int64-7":  -2147483648,
		"A-int64-8":  0,
		"A-int8-1":   127,
		"A-int8-2":   127,
		"A-int8-3":   128,
		"A-int8-4":   128,
		"A-int8-5":   -128,
		"A-int8-6":   -128,
		"A-int8-7":   -129,
		"A-int8-8":   -129,
		"A-uint16-1": 65535,
		"A-uint16-2": 65535,
		"A-uint16-3": 65536,
		"A-uint16-4": 65536,
		"A-uint32-1": -2147483648,
		"A-uint32-2": -1,
		"A-uint32-3": -2147483648,
		"A-uint32-4": 0,
		"A-uint64-1": -2147483648,
		"A-uint64-2": -1,
		"A-uint64-3": -2147483648,
		"A-uint64-4": -1,
		"A-uint8-1":  255,
		"A-uint8-2":  255,
		"A-uint8-3":  256,
		"A-uint8-4":  256,
	}
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-getint.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	}
	for k, v := range testMap {
		v1, err := GetInt32(k)
		if err != nil {
			t.Errorf("GetInt32 key=%s failed:%v", k, err)
		} else {
			if v1 != v {
				t.Errorf("testKey=%v get=%v testExpect=%v failed", k, v1, v)
			}
		}
	}
	_, err = GetInt("A-test-string-invalid")
	if err == nil {
		t.Errorf("GetInt32 key=A-test-string-invalid failed")
	}

	v1, err := GetInt("A-test-bool-true")
	if err != nil {
		t.Errorf("GetInt32 key=A-test-bool-true failed:%v", err)
	} else {
		if v1 != 1 {
			t.Errorf("GetInt32 key=A-test-bool-true failed:expect 1,but get v1=%v", v1)
		}
	}
	v2, err := GetInt("A-test-bool-false")
	if err != nil {
		t.Errorf("GetInt32 key=A-test-bool-false failed:%v", err)
	} else {
		if v2 != 0 {
			t.Errorf("GetInt32 key=A-test-bool-false failed:expect 0,but get v1=%v", v1)
		}
	}
}

func TestGetInt64(t *testing.T) {
	testMap := map[string]int64{
		"A-int16-1":  32767,
		"A-int16-2":  32767,
		"A-int16-3":  32768,
		"A-int16-4":  32768,
		"A-int16-5":  -32768,
		"A-int16-6":  -32768,
		"A-int16-7":  -32769,
		"A-int16-8":  -32769,
		"A-int32-1":  2147483647,
		"A-int32-2":  2147483647,
		"A-int32-3":  2147483648,
		"A-int32-4":  2147483648,
		"A-int32-5":  -2147483648,
		"A-int32-6":  -2147483648,
		"A-int32-7":  -2147483649,
		"A-int32-8":  -2147483649,
		"A-int64-1":  -9223372036854775808,
		"A-int64-2":  9223372036854775807,
		"A-int64-3":  -9223372036854775808,
		"A-int64-4":  9223372036854775807,
		"A-int64-5":  -9223372036854775808,
		"A-int64-6":  -9223372036854775808,
		"A-int64-7":  -9223372036854775808,
		"A-int64-8":  -9223372036854775808,
		"A-int8-1":   127,
		"A-int8-2":   127,
		"A-int8-3":   128,
		"A-int8-4":   128,
		"A-int8-5":   -128,
		"A-int8-6":   -128,
		"A-int8-7":   -129,
		"A-int8-8":   -129,
		"A-uint16-1": 65535,
		"A-uint16-2": 65535,
		"A-uint16-3": 65536,
		"A-uint16-4": 65536,
		"A-uint32-1": 4294967295,
		"A-uint32-2": 4294967295,
		"A-uint32-3": 4294967296,
		"A-uint32-4": 4294967296,
		"A-uint64-1": -9223372036854775808,
		"A-uint64-2": 9223372036854775807,
		"A-uint64-3": -9223372036854775808,
		"A-uint64-4": 9223372036854775807,
		"A-uint8-1":  255,
		"A-uint8-2":  255,
		"A-uint8-3":  256,
		"A-uint8-4":  256,
	}
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-getint.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	}
	for k, v := range testMap {
		v1, err := GetInt64(k)
		if err != nil {
			t.Errorf("GetInt64 key=%s failed:%v", k, err)
		} else {
			if v1 != v {
				t.Errorf("testKey=%v get=%v testExpect=%v failed", k, v1, v)
			}
		}
	}
	_, err = GetInt64("A-test-string-invalid")
	if err == nil {
		t.Errorf("GetInt64 key=A-test-string-invalid failed")
	}

	v1, err := GetInt64("A-test-bool-true")
	if err != nil {
		t.Errorf("GetInt64 key=A-test-bool-true failed:%v", err)
	} else {
		if v1 != 1 {
			t.Errorf("GetInt64 key=A-test-bool-true failed:expect 1,but get v1=%v", v1)
		}
	}
	v2, err := GetInt64("A-test-bool-false")
	if err != nil {
		t.Errorf("GetInt64 key=A-test-bool-false failed:%v", err)
	} else {
		if v2 != 0 {
			t.Errorf("GetInt64 key=A-test-bool-false failed:expect 0,but get v1=%v", v1)
		}
	}
}

func TestGetUint32(t *testing.T) {
	testMap := map[string]uint32{
		"A-int16-1":  32767,
		"A-int16-2":  32767,
		"A-int16-3":  32768,
		"A-int16-4":  32768,
		"A-int16-5":  32768,
		"A-int16-6":  0,
		"A-int16-7":  0,
		"A-int16-8":  0,
		"A-int32-1":  2147483647,
		"A-int32-2":  2147483647,
		"A-int32-3":  2147483648,
		"A-int32-4":  2147483648,
		"A-int32-5":  2147483648,
		"A-int32-6":  2147483648,
		"A-int32-7":  2147483649,
		"A-int32-8":  0,
		"A-int64-1":  2147483649,
		"A-int64-2":  2147483649,
		"A-int64-3":  2147483649,
		"A-int64-4":  2147483649,
		"A-int64-5":  0,
		"A-int64-6":  0,
		"A-int64-7":  2147483649,
		"A-int64-8":  2147483649,
		"A-int8-1":   127,
		"A-int8-2":   127,
		"A-int8-3":   128,
		"A-int8-4":   128,
		"A-int8-5":   128,
		"A-int8-6":   128,
		"A-int8-7":   129,
		"A-int8-8":   129,
		"A-uint16-1": 65535,
		"A-uint16-2": 65535,
		"A-uint16-3": 65536,
		"A-uint16-4": 65536,
		"A-uint32-1": 4294967295,
		"A-uint32-2": 4294967295,
		"A-uint32-3": 4294967295,
		"A-uint32-4": 4294967295,
		"A-uint64-1": 4294967295,
		"A-uint64-2": 0,
		"A-uint64-3": 4294967295,
		"A-uint64-4": 4294967295,
		"A-uint8-1":  255,
		"A-uint8-2":  255,
		"A-uint8-3":  256,
		"A-uint8-4":  256,
	}
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-getint.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	}
	for k, v := range testMap {
		v1, err := GetUint32(k)
		if err != nil {
			if v != 0 {
				t.Errorf("GetUint32 key=%s failed:%v", k, err)
			}
		} else {
			if v1 != v {
				t.Errorf("testKey=%v get=%v testExpect=%v failed", k, v1, v)
			}
		}
	}
	_, err = GetUint32("A-test-string-invalid")
	if err == nil {
		t.Errorf("GetUint32 key=A-test-string-invalid failed")
	}

	v1, err := GetUint32("A-test-bool-true")
	if err != nil {
		t.Errorf("GetUint32 key=A-test-bool-true failed:%v", err)
	} else {
		if v1 != 1 {
			t.Errorf("GetUint32 key=A-test-bool-true failed:expect 1,but get v1=%v", v1)
		}
	}
	v2, err := GetUint32("A-test-bool-false")
	if err != nil {
		t.Errorf("GetUint32 key=A-test-bool-false failed:%v", err)
	} else {
		if v2 != 0 {
			t.Errorf("GetUint32 key=A-test-bool-false failed:expect 0,but get v1=%v", v1)
		}
	}
}

func TestFloat64(t *testing.T) {
	testMap := map[string]float64{
		"A": 1.1,
		"B": 1.11,
		"C": 1.11,
		"D": 1,
		"E": 1,
		"F": 1,
		"G": 0,
	}
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-getfloat64.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	} else {
		for k, v := range testMap {
			v1, err := GetFloat64(k)
			if err != nil {
				if v != 0 {
					t.Errorf("GetFloat64 key=%s failed:%v", k, err)
				}
			} else {
				if v1 != v {
					t.Errorf("testKey=%v get=%v testExpect=%v failed", k, v1, v)
				}
			}
		}
		_, err = GetFloat64("J")
		if err == nil {
			t.Errorf("GetFloat64 key=J failed")
		}
	}
}

func TestGetIntSlice(t *testing.T) {
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-getintslice.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	} else {
		v, err := GetIntSlice("A")
		if err != nil {
			t.Error("GetIntSlice failed:", err)
		} else {
			if len(v) != 6 {
				t.Errorf("GetIntSlice invalid:%v", v)
			}
		}
	}
}

func TestSubGet(t *testing.T) {
	testKeySuffixMap := []string{
		"obj", "int", "string", "bool", "float",
	}
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-getsub.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	} else {
		allKeys := make([]string, 0)

		for i := 1; i <= 7; i++ {
			totalKeyB := ""
			curkeyB := ""
			totalKey := ""
			curkey := ""
			for kIndex := 1; kIndex <= i; kIndex++ {
				curkey += "A"
				curkeyB += "B"
			}
			if i > 1 {
				preKey := ""
				for j := 1; j < i; j++ {
					for kIndex := 1; kIndex <= j; kIndex++ {
						preKey += "A"
					}
					if j < i {
						preKey += "."
					}
				}
				totalKey = preKey + curkey
				totalKeyB = preKey + curkeyB
			} else {
				totalKey = curkey
				totalKeyB = curkeyB
			}

			// A
			// A.AA
			// A.AA.AAA
			// for kIndex := 1; kIndex <= i; kIndex++ {
			// 	key += "A"
			// }

			for _, suf := range testKeySuffixMap {
				newkey := totalKey + "-" + suf
				newkeyB := totalKeyB + "-" + suf
				allKeys = append(allKeys, newkey)
				allKeys = append(allKeys, newkeyB)
			}
		}
		for _, k := range allKeys {
			tmps := strings.Split(k, "-")
			if tmps[len(tmps)-1] == "bool" {
				_, err := GetBool(k)
				if err != nil {
					t.Errorf("key=%s get bool failed:%v\n", k, err)
				}
			} else if tmps[len(tmps)-1] == "int" {
				_, err := GetInt(k)
				if err != nil {
					t.Errorf("key=%s get int failed:%v\n", k, err)
				}
			} else if tmps[len(tmps)-1] == "float" {
				_, err := GetFloat64(k)
				if err != nil {
					t.Errorf("key=%s get float failed:%v\n", k, err)
				}
			} else if tmps[len(tmps)-1] == "string" {
				_, err := GetString(k)
				if err != nil {
					t.Errorf("key=%s get string failed:%v\n", k, err)
				}
			}

		}
	}
}
