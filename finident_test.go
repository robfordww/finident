package finident

import "testing"

func TestValidateLEICodes(t *testing.T) {
	expectFail(t, "lei")
	expectFail(t, "NO123451234121212312")
	expectFail(t, "NO1$0051234121212312")
	expectFail(t, "NO1q00512341212Ø212")
	expectFail(t, "NO120051234121212312")
	expectFail(t, "NO120051234121212332")
	expectFail(t, "NO120051234121212333")
	expectOK(t, "5493004W1IPC50878Z34")
	expectOK(t, "815600D0B9FB2B70AA10")
	expectOK(t, "5493008WYXIP9CE4ER31")
}

func TestValidationOfMod97(t *testing.T) {
	if !Validatemod97("000100001234567890194252950") {
		t.Fail()
	}
	if !Validatemod97("123443211234567890172") {
		t.Fail()
	}
	if Validatemod97("123443211234567890173") {
		t.Fail()
	}
}

func TestCalculation(t *testing.T) {
	if c := CalculateChecksum("1234432112345678901"); c != "72" {
		t.Errorf("%v != 72", c)
	}
	if c := CalculateChecksum("5493004W1IPC50878Z"); c != "34" {
		t.Errorf("%v != 34", c)
	}
}

func BenchmarkLeiValidation(b *testing.B) {
	lei := "5493004W1IPC50878Z34"
	for i := 0; i < b.N; i++ {
		//b.StartTimer()
		ValidateLEI(lei)
		//b.StopTimer()
		lei = lei[19:] + lei[:19] // Rotate
	}
}

func expectFail(t *testing.T, lei string) {
	r, e := ValidateLEI(lei)
	if r == true || e == nil {
		t.Errorf("LEI: '%s' is valid", lei)
	}
}

func expectOK(t *testing.T, lei string) {
	r, e := ValidateLEI(lei)
	if r != true || e != nil {
		t.Errorf("Lei: '%s' failed with error %s", lei, e)
	}
}
