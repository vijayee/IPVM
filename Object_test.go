package ipvm

import (
	"testing"
)

type testData struct {
	input  interface{}
	output interface{}
}

var testSetData map[string]testData, objArray map[string]Object

func TestSetObject(t *testing.T) {
	for tk, d := range testSetData {
		testObj := new(Object)
		err := testObj.Set(d.input)
		if err != nil {
			t.Error("Error occurred on ", tk, ": ", err.Error())
		}
	}

}
func TestGetObject(t *testing.T) {
	for tk, d := range testSetData {
		testObj := new(Object)
		err := testObj.Set(d.input)
		if err != nil {
			t.Error("Error occurred on ", tk, ": ", err.Error())
		}
	}

}
func TestMain(m *testing.M) {
	testSetData = make(map[string]testData)
	objArray = make(map[string]Object)
	testSetData["stringTest"] = testData{"something", nil}
	testSetData["boolTest"] = testData{false, nil}
}
