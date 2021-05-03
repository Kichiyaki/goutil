package strutil

import "testing"

func TestPascalCase(t *testing.T) {
	tests := []test{
		{
			data:           "t_e_s_t",
			expectedResult: "TEST",
		},
		{
			data:           "test",
			expectedResult: "Test",
		},
		{
			data:           "test_test",
			expectedResult: "TestTest",
		},
		{
			data:           "test_id",
			expectedResult: "TestId",
		},
		{
			data:           "test.id",
			expectedResult: "TestId",
		},
		{
			data:           "test.id_test",
			expectedResult: "TestIdTest",
		},
	}
	for _, singleTest := range tests {
		result := PascalCase(singleTest.data, '_', '.')
		if result != singleTest.expectedResult {
			t.Errorf("expected %s, got %s", singleTest.expectedResult, result)
		}
	}
}
