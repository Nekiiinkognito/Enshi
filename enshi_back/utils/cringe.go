package utils

import (
	"fmt"
	"reflect"
)

func ConvertInt64ToStringInStruct(input any) (any, error) {
	origVal := reflect.ValueOf(input)

	// Ensure input is a struct or pointer to a struct
	if origVal.Kind() == reflect.Ptr {
		origVal = origVal.Elem()
	}
	if origVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or a pointer to a struct")
	}

	// Create a new instance of the same type
	newStruct := reflect.New(origVal.Type()).Elem()

	// Iterate through fields
	for i := 0; i < origVal.NumField(); i++ {
		field := origVal.Field(i)
		newField := newStruct.Field(i)

		if !newField.CanSet() {
			// Skip unexported fields
			continue
		}

		switch field.Kind() {
		case reflect.Int64:
			// Convert int64 fields to string if the target is compatible
			// if newField.Kind() == reflect.Int64 {
			// 	fmt.Print("aqwrqfwq", field)
			// newField.Set(strconv.FormatInt(field.Int(), 10))
			// }
			newField.SetString("asd")
		default:
			// Copy other fields directly
			if newField.Type() == field.Type() {
				newField.Set(field)
			}
		}
	}

	return newStruct.Interface(), nil
}
