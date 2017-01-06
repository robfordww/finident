package finident

import "testing"

func TestISIN(t *testing.T) {
	invalid := []string{"GB00B0SWJX3", "1S25152CMN38", "GB00B0SWJX34ZA", "GB00B0/JX4ZA", "GB00B0SWJX35"}
	for _, i := range invalid {
		if b, err := ValidateISIN(i); b == true {
			t.Errorf("Expected %s to fail, but it passed: %v", i, err)
		}
	}

	valid := []string{"GB00B0SWJX34", "US25152CMN38", "US0378331005"}
	for _, i := range valid {
		if b, err := ValidateISIN(i); b != true {
			t.Errorf("Expected %s to validate, but it didn't: %v", i, err)
		}
	}
}

func BenchmarkISINValidation(b *testing.B) {
	valid := []string{"GB00B0SWJX34",
		"US25152CMN38", "US0378331005", "LI0123534161", "LI0123534146", "DE0009750026",
		"AT0000A0GWN4", "AT0000821095", "AT0000708367", "DE000A0M80H2", "AT0000A0HQY1",
		"AT0000A0V5U6", "AT0000824701", "AT0000855820", "AT0000622923", "AT0000A07RY4",
		"AT0000736392", "AT0000A07RZ1", "AT0000793732", "AT0000835681", "AT0000A0HQX3",
		"AT0000855846", "AT0000855861", "AT0000858204", "AT0000A0HR07", "AT0000A00LF1",
		"AT0000A07RW8", "AT0000A07RX6"}

	/*
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				if r, e :=
					ValidateISIN(valid[i%len(valid)]); r != true {
					b.Errorf("Validation failed. %v", e)
					b.FailNow()
				}
				i++
			}
		})*/

	for i := 0; i < b.N; i++ {
		if r, e :=
			ValidateISIN(valid[i%len(valid)]); r != true {
			b.Errorf("Validation failed. %v", e)
			b.FailNow()
		}
	}
}

// LEI code tests
func TestValidateLEICodes(t *testing.T) {
	expectFail(t, "lei")
	expectFail(t, "NO123451234121212312")
	expectFail(t, "NO1$0051234121212312")
	expectFail(t, "NO1q00512341212Ø212")
	expectFail(t, "NO120051234121212312")
	expectFail(t, "NO120051234121212332")
	expectFail(t, "NO120051234121212333")
	expectFail(t, "NO120051234121212Ø3")
	expectFail(t, "5493004W1IPC50878Z35")
	expectFail(t, "5493004W1IPC5087Ø34")
	expectFail(t, "5493004W1IPC5/878Z34")
	expectOK(t, "5493004W1IPC50878Z34")
	expectOK(t, "815600D0B9FB2B70AA10")
	expectOK(t, "5493008WYXIP9CE4ER31")
}

func TestValidationOfMod97(t *testing.T) {
	if !Validatemod97("000100001234567890194252950") {
		t.Fail()
	}
	if Validatemod97("00010000123456789019425295C") {
		t.Fail()
	}
	if Validatemod97("0001000012345678901942529B0") {
		t.Fail()
	}
	if !Validatemod97("815600a0B9FB2B70AA10") {
		t.Fail()
	}
}

func TestValidationOfMod97Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Validatemod97("12344321123456789013/")
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
	lei := []string{"5493004W1IPC50878Z34", "529900H1R1AQB4PG9763",
		"5299004VX6J1Y7H52S80", "549300FXBIWWGK7T0Y98", "EYG9EUSWTIUWOF7QFT34",
		"529900YXKZPO3Y7GWS93", "5299005LGIYUBJ86DT56", "529900A4VLNE8WSFTX76",
		"529900JSN7UYZYMMO265", "529900VN54ULT9WBKE58", "815600D7B9CC3A5B7344",
		"8156008E0560EE7C0151", "815600605A5B1EA19986", "8RS0AKOLN987042F2V04",
		"815600D05CA3A663CE35", "5299006S3ALB1X1PU159", "35GDVHRBMFE7NWATNM84",
		"213800IWGUQS3U4V8953", "549300SQ4ZSVSWC6H750", "213800PQLSKZ25LSII39",
		"LUZQVYP4VS22CLWDAR65", "529900VVQ4470YJ67K26", "5493005LM11U105HR746",
		"222100BLL26OLIPJ3F50", "2138003QX1RSCHWUB420", "213800G95T751RN2CT94"}

	for i := 0; i < b.N; i++ {
		//b.StartTimer()
		ValidateLEI(lei[i%len(lei)])
		//b.StopTimer()
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

func TestSum(t *testing.T) {
	if a, b := sumOfDigits(-123)*-1, sumOfDigits(123); a != b {
		t.Errorf("sumOfDigits(-123)*-1[%v] != sumOfDigits(123)[%v]", a, b)
	}
	if sumOfDigits(123) != 6 {
		t.Error("sumOfDigits(123) != 6")
	}
	if sumOfDigits(0) != 0 {
		t.Error("sumOfDigits(0) != 0")
	}
}
