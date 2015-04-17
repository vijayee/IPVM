package ipvm

import (
	vm "github.com/robertkrimen/otto"
)

//simplified javascript object that unifies functions of Object and Value of otto
type Object struct {
	value vm.Value
}

//define an interface for which an api may include objects consistently
type Objecter interface {
	Object() Object
}

//return the interface of the value
func (o *Object) Get() (interface{}, error) {
	return o.value.Export()
}

//set an interface to be the object's Value
func (o *Object) Set(value interface{}) error {
	var err error
	o.value, err = vm.ToValue(value)
	return err
}

//get an array of the properties the object contains
func (o *Object) Properties(value interface{}) []string {
	return o.value.Object().Keys()
}

//set an object's properties
func (o *Object) SetProp(name string, value interface{}) error {
	return o.value.Object().Set(name, value)
}

//get an object's properties
func (o *Object) GetProp(name string) (Object, error) {
	value, err := o.value.Object().Get(name)
	return Object{value}, err
}

//return type of the object
// The return value will (generally) be one of:
//
//		Object
//		Function
//		Array
//		String
//		Number
//		Boolean
//		Date
//		RegExp
//
func (o *Object) Type() string {
	return o.value.Object().Class()
} //return boolean

func (o *Object) IsBoolean() bool {
	return o.value.IsBoolean()

}

func (o *Object) IsDefined() bool {
	return o.value.IsDefined()
}

// IsFunction will return true if value is a function.
func (o *Object) IsFunction() bool {
	return o.value.IsFunction()
}

// IsNaN will return true if value is NaN (or would convert to NaN).
func (o *Object) IsNaN() bool {
	return o.value.IsNaN()
}

// IsNull will return true if the value is null, and false otherwise.
func (o *Object) IsNull() bool {
	return o.value.IsNull()
}

// IsNumber will return true if value is a number (primitive).
func (o *Object) IsNumber() bool {
	return o.value.IsNumber()
}

// IsObject will return true if value is an object.
func (o *Object) IsObject() bool {
	return o.value.IsObject()
}

// IsPrimitive will return true if value is a primitive (any kind of primitive)
func (o *Object) IsPrimitive() bool {
	return o.value.IsPrimitive()
}

// IsString will return true if value is a string (primitive).
func (o *Object) IsString() bool {
	return o.value.IsString()
}

// IsUndefined will return true if the value is undefined, and false otherwise.
func (o *Object) IsUndefined() bool {
	return o.value.IsUndefined()
}

//return a string if the object can be represented as such
func (o *Object) String() string {
	return o.value.String()
}
