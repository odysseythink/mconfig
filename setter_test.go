package mconfig

import (
	"testing"

	"mlib.com/mlog"
)

func TestBaseSet(t *testing.T) {
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-baseset.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	} else {
		v1, err := GetInt("A")
		if err != nil {
			t.Error("GetInt A failed:", err)
		} else {
			err = Set("A", v1+2)
			if err != nil {
				t.Error("Set A failed:", err)
			} else {
				v2, err := GetInt("A")
				if err != nil {
					t.Error("after set a,GetInt A failed:", err)
				} else {
					if v2 != v1+2 {
						t.Error("set a failed:invalid set value=", v2)
					} else {
						err = WriteConfig()
						if err != nil {
							t.Error("WriteConfig failed:", err)
						} else { // test isChanged flag can work or not
							err = Set("A", v1+2)
						}

					}
				}
			}
		}
	}
}

func TestSetIsChanged(t *testing.T) {
	mlog.SetLogDir("./logs")
	SetConfigFile("./test-baseset.json")
	err := ReadInConfig()
	if err != nil {
		t.Error("ReadInConfig failed:", err)
	} else {
		v1, err := GetInt("A")
		if err != nil {
			t.Error("GetInt A failed:", err)
		} else {
			err = Set("A", v1)
			if err != nil {
				t.Error("Set A failed:", err)
			} else {
				err = WriteConfig()
				if err != nil {
					t.Error("WriteConfig failed:", err)
				}
			}
		}
	}
}
