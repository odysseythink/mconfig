package mconfig

// func TestBasics(t *testing.T) {
// 	mlog.SetLogDir("./logs")
// 	SetConfigFile("./config.json")
// 	err := ReadInConfig()
// 	if err != nil {
// 		t.Error("ReadInConfig failed:", err)
// 	}
// }

// func TestLevel1SetAndWriter(t *testing.T) {
// 	mlog.SetLogDir("./logs")
// 	SetConfigFile("./config.json")
// 	err := ReadInConfig()
// 	if err != nil {
// 		t.Error("ReadInConfig failed:", err)
// 	}
// 	t.Log("replace test start--------------------------")
// 	err = Set("A", "replace A, with string")
// 	if err != nil {
// 		t.Error("replace A, with string failed:", err)
// 	}
// 	if GetString("A") != "replace A, with string" {
// 		t.Errorf("key A set string failed")
// 	}
// 	err = Set("B", true)
// 	if err != nil {
// 		t.Error("replace B, with bool failed:", err)
// 	}
// 	if GetBool("B") != true {
// 		t.Errorf("key B set bool failed")
// 	}
// 	err = Set("C", 3.14)
// 	if err != nil {
// 		t.Error("replace C, with float failed:", err)
// 	}
// 	if GetFloat64("C") != 3.14 {
// 		t.Errorf("key C set float failed")
// 	}
// 	err = Set("D", 3)
// 	if err != nil {
// 		t.Error("replace D, with float failed:", err)
// 	}
// 	if GetInt("D") != 3 {
// 		t.Errorf("key D set float failed")
// 	}
// 	t.Log("replace test end----------------------------")
// 	WriteConfig()
// }

// func TestLevel2SetAndWriter(t *testing.T) {
// 	mlog.SetLogDir("./logs")
// 	SetConfigFile("./config.json")
// 	err := ReadInConfig()
// 	if err != nil {
// 		t.Error("ReadInConfig failed:", err)
// 	}
// 	t.Log("replace level2 exist key test start--------------------------")
// 	err = Set("E.A", "replace E.A, with string")
// 	if err != nil {
// 		t.Error("replace E.A, with string failed:", err)
// 	}
// 	if GetString("E.A") != "replace E.A, with string" {
// 		t.Errorf("key E.A set string failed")
// 	}
// 	err = Set("E.B", true)
// 	if err != nil {
// 		t.Error("replace E.B, with bool failed:", err)
// 	}
// 	if GetBool("E.B") != true {
// 		t.Errorf("key E.B set bool failed")
// 	}
// 	err = Set("E.C", 3.14)
// 	if err != nil {
// 		t.Error("replace E.C, with float failed:", err)
// 	}
// 	if GetFloat64("E.C") != 3.14 {
// 		t.Errorf("key C set float failed")
// 	}
// 	err = Set("E.D", 3)
// 	if err != nil {
// 		t.Error("replace E.D, with float failed:", err)
// 	}
// 	if GetInt("E.D") != 3 {
// 		t.Errorf("key E.D set float failed")
// 	}
// 	t.Log("replace level2 exist key test end----------------------------")
// 	WriteConfig()
// }

// func TestLevel3SetAndWriter(t *testing.T) {
// 	mlog.SetLogDir("./logs")
// 	SetConfigFile("./config.json")
// 	err := ReadInConfig()
// 	if err != nil {
// 		t.Error("ReadInConfig failed:", err)
// 	}
// 	t.Log("replace level3 exist key test start--------------------------")
// 	err = Set("E.E.A", "replace E.E.A, with string")
// 	if err != nil {
// 		t.Error("replace E.E.A, with string failed:", err)
// 	}
// 	if GetString("E.E.A") != "replace E.E.A, with string" {
// 		t.Errorf("key E.E.A set string failed")
// 	}
// 	err = Set("E.E.B", true)
// 	if err != nil {
// 		t.Error("replace E.E.B, with bool failed:", err)
// 	}
// 	if GetBool("E.E.B") != true {
// 		t.Errorf("key E.E.B set bool failed")
// 	}
// 	err = Set("E.E.C", 3.14)
// 	if err != nil {
// 		t.Error("replace E.E.C, with float failed:", err)
// 	}
// 	if GetFloat64("E.E.C") != 3.14 {
// 		t.Errorf("key E.E.C set float failed")
// 	}
// 	err = Set("E.E.D", 3)
// 	if err != nil {
// 		t.Error("replace E.E.D, with float failed:", err)
// 	}
// 	if GetInt("E.E.D") != 3 {
// 		t.Errorf("key E.E.D set float failed")
// 	}
// 	t.Log("replace level3 exist key test end----------------------------")
// 	WriteConfig()
// }
