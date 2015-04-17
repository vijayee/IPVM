package ipvm

import (
	"fmt"
	"testing"
)

type testData struct {
	input  interface{}
	output interface{}
}

var testSetData map[string]testData
var objArray map[string]*Object

func TestSetObject(t *testing.T) {
	for tk, d := range testSetData {
		objArray[tk] = new(Object)
		err := objArray[tk].Set(d.input)
		if err != nil {
			t.Errorf("Set() Error occurred on %s: %s", tk, err.Error())

		}
		fmt.Printf("%s was Set() succesfully \n", tk)
	}

}
func TestGetObject(t *testing.T) {
	for tk, d := range objArray {
		_, err := d.Get()
		if err != nil {
			t.Errorf("Get() Error occurred on %s: %s", tk, err.Error())
		}
		fmt.Printf("%s was Get() succesfully  of object %s\n", tk)
	}

}
func TestMain(m *testing.M) {
	testSetData = make(map[string]testData)
	objArray = make(map[string]*Object)
	testSetData["stringTest"] = testData{"something", "something"}
	testSetData["boolTest"] = testData{false, false}
}
