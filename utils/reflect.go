package utils

import "reflect"

func SetVal(obj interface{}, fieldName string, value interface{}) bool {

	// Get the underlying reflect.Value for the struct
	reflectObj := reflect.ValueOf(obj)

	// Dereference pointer if needed
	if reflectObj.Kind() == reflect.Ptr {
		reflectObj = reflectObj.Elem()
	}

	// Get the field value
	fieldVal := reflectObj.FieldByName(fieldName)

	// If it's settable, set the value
	if fieldVal.CanSet() {
		fieldVal.Set(reflect.ValueOf(value))
		return true
	}
	return false
}
