package envutil

import (
	"os"
	"testing"
)

const key = "TEMP"

type test struct {
	data           string
	expectedResult interface{}
}

func TestGetenvInt(t *testing.T) {
	tests := []test{
		{
			data:           "15",
			expectedResult: 15,
		},
		{
			data:           "-15",
			expectedResult: -15,
		},
		{
			data:           "63211231521321321",
			expectedResult: 63211231521321321,
		},
		{
			data:           "asd",
			expectedResult: 0,
		},
	}
	for _, singleTest := range tests {
		if err := os.Setenv(key, singleTest.data); err != nil {
			t.Errorf("Couldn't set env variable: %s", err)
		}
		result := GetenvInt(key)
		if singleTest.expectedResult != result {
			t.Errorf("Expected %d, got %d", singleTest.expectedResult, result)
		}
	}
}

func TestGetenvBool(t *testing.T) {
	tests := []test{
		{
			data:           "1",
			expectedResult: true,
		},
		{
			data:           "true",
			expectedResult: true,
		},
		{
			data:           "asd",
			expectedResult: false,
		},
		{
			data:           "false",
			expectedResult: false,
		},
	}
	for _, singleTest := range tests {
		if err := os.Setenv(key, singleTest.data); err != nil {
			t.Errorf("Couldn't set env variable: %s", err)
		}
		result := GetenvBool(key)
		if singleTest.expectedResult != result {
			t.Errorf("Expected %v, got %v", singleTest.expectedResult, result)
		}
	}
}
