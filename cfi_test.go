package finident

import (
	"testing"
)

func TestCFIs(t *testing.T) {
	okList := []string{"ESVTOB", "EFRRCN", "DDBVGF", "DMBXXR"}
	failList := []string{"ESVTOZ", "EFRRECN", "ZDBVGF", "CPBXXR"}
	for i := range okList {
		if !IsValidCFI(okList[i]) {
			t.Error(okList[i])
		}
	}
	for i := range failList {
		if IsValidCFI(failList[i]) {
			t.Error(okList[i])
		}
	}
}

func TestGeneration(t *testing.T) {
	res := GenCFICombinations()
	for i := range res {
		if !IsValidCFI(res[i]) {
			t.Fail()
		}
	}
	t.Log("Count", len(res))
	if len(res) < 1000 {
		t.Fail()
	}
	t.Log(res[0], res[999])
}
