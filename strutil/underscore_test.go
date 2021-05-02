package strutil

import "testing"

func TestUnderscore(t *testing.T) {
	tests := []test{
		{
			data:           "xYz",
			expectedResult: "x_yz",
		},
		{
			data:           "TEST",
			expectedResult: "test",
		},
		{
			data:           "TeST",
			expectedResult: "te_st",
		},
		{
			data:           "PascalCase",
			expectedResult: "pascal_case",
		},
		{
			data:           "nothing_to_change",
			expectedResult: "nothing_to_change",
		},
		{
			data:           "camelCase",
			expectedResult: "camel_case",
		},
	}
	for _, singleTest := range tests {
		result := Underscore(singleTest.data)
		if result != singleTest.expectedResult {
			t.Errorf("expected %s, got %s", singleTest.expectedResult, result)
		}
	}
}
