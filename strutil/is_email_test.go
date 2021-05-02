package strutil

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []test{
		{
			data:           "test@gmail.com",
			expectedResult: true,
		},
		{
			data:           "",
			expectedResult: false,
		},
		{
			data:           "test",
			expectedResult: false,
		},
		{
			data:           "test@asdad1231512asdookolqwq.com",
			expectedResult: false,
		},
		{
			data:           "test@ovh.com",
			expectedResult: true,
		},
	}
	for _, singleTest := range tests {
		result := IsEmail(singleTest.data)
		if result != singleTest.expectedResult {
			t.Errorf("expected %v, got %v", singleTest.expectedResult, result)
		}
	}
}
