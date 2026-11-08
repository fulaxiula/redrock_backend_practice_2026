package even

import "testing"

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Errorf("IsEven(2) = false, want true")
	}
	if IsEven(3) {
		t.Errorf("IsEven(3) = true, want false")
	}
	if !IsEven(0) {
		t.Errorf("IsEven(0) = false, want true")
	}
	if !IsEven(-2) {
		t.Errorf("IsEven(-2) = false, want true")
	}
}