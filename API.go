package ipvm

import (
	"errors"
	"reflect"
)

var api map[string]interface{}

func init() {
	if api == nil {
		api = make(map[string]interface{})
	}
}

//Define an API to referenced by a name for a contract
// You may only define an api that is of following
//		Object
//		func(Object) Object
//		func (Object)
func Define(name string, object interface{}) error {
	switch t := object.(type) {
	case Object:
		api[name] = object
		return nil
	case Objecter:
		api[name] = t.Object()
		return nil
	default:
		// Investigate whether it is a function of the aforementioned signatures
		typ := reflect.TypeOf(object)
		//get the reflection type of the Object data type
		objectType := reflect.TypeOf((*Object)(nil)).Elem()
		switch typ.Kind() {
		case reflect.Func:
			switch {
			case typ.NumOut() > 1:
				return errors.New("Invalid IPVM Object")
			case typ.NumOut() > 0 && typ.Out(0) != objectType:
				return errors.New("Invalid IPVM Object")
			default:
				for i := 0; i < typ.NumIn(); i++ {
					if typ.In(i) != objectType {
						return errors.New("Invalid IPVM Object")
					}
				}
				api[name] = object
				return nil
			}

		default:
			return errors.New("Invalid IPVM Object")
		}

	}
}
