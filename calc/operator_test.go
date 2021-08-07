package calc

import "testing"

func TestPickOp(t *testing.T) {
	if _, ok := pickOp("+"); !ok {
		t.Log("+ is op")
		t.Fail()
	}
	if _, ok := pickOp("-"); !ok {
		t.Log("- is op")
		t.Fail()
	}
	if _, ok := pickOp("*"); !ok {
		t.Log("* is op")
		t.Fail()
	}
	if _, ok := pickOp("/"); !ok {
		t.Log("/ is op")
		t.Fail()
	}
	if _, ok := pickOp("00"); ok {
		t.Log("00 is not op")
		t.Fail()
	}
	if _, ok := pickOp("++"); ok {
		t.Log("++ is not op")
		t.Fail()
	}
}

func TestPickOpBsPre(t *testing.T) {
	if pickOpBsPre([]byte("+")) == nil {
		t.Log("+ is op")
		t.Fail()
	}
	if pickOpBsPre([]byte("-")) == nil {
		t.Log("- is op")
		t.Fail()
	}
	if pickOpBsPre([]byte("*")) == nil {
		t.Log("* is op")
		t.Fail()
	}
	if pickOpBsPre([]byte("/")) == nil {
		t.Log("/ is op")
		t.Fail()
	}
	if pickOpBsPre([]byte("++")) == nil {
		t.Log("++ has prefix +")
		t.Fail()
	}
	if pickOpBsPre([]byte("+0")) == nil {
		t.Log("+0 has prefix +")
		t.Fail()
	}
	if pickOpBsPre([]byte("0")) != nil {
		t.Log("0 is not op")
		t.Fail()
	}
}
