package strutil

import "testing"

func TestMinify(t *testing.T) {
	tests := []test{
		{
			data: `
test

test
`,
			expectedResult: "test test",
		},
		{
			data: `
test

test

			test			test
`,
			expectedResult: "test test test test",
		},
	}
	for _, singleTest := range tests {
		result := Minify(singleTest.data, " ")
		if result != singleTest.expectedResult {
			t.Errorf("expected %v, got %v", singleTest.expectedResult, result)
		}
	}
}
