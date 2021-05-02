package safeptr

import "testing"

func TestSafeBoolPointer(t *testing.T) {
	if !SafeBoolPointer(nil, true) {
		t.Errorf("expected true, got false")
	}
	val := true
	if !SafeBoolPointer(&val, true) {
		t.Errorf("expected true, got false")
	}
}

func TestSafeStringPointer(t *testing.T) {
	if result := SafeStringPointer(nil, ""); result != "" {
		t.Errorf("expected empty string, got %v", result)
	}
	val := "asd"
	if result := SafeStringPointer(&val, ""); result != val {
		t.Errorf("expected %s, got %v", val, result)
	}
}

func TestSafeIntPointer(t *testing.T) {
	if result := SafeIntPointer(nil, 0); result != 0 {
		t.Errorf("expected 0, got %v", result)
	}
	val := 15
	if result := SafeIntPointer(&val, 0); result != val {
		t.Errorf("expected %d, got %v", val, result)
	}
}
