package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Decode a an array of bytes and return a reflected value of a struct
func Decode(reg map[string]interface{}, d []byte, k string) (reflect.Value, error) {
	t := reflect.TypeOf(reg[k])
	pv := reflect.New(t.Elem())
	err := json.Unmarshal(d, pv.Interface())
	return pv.Elem(), err
}

// Set the field of a reflected value of struct
func SetField(v reflect.Value, fn string, fv interface{}) error {
	fmt.Printf("%+v\n", v.Type())
	if !v.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}
	findJsonName := func(t reflect.StructTag) (string, error) {
		if jt, ok := t.Lookup("json"); ok {
			return strings.Split(jt, ",")[0], nil
		}
		return "", fmt.Errorf("tag provided does not define a json tag %s", fn)
	}
	fieldNames := map[string]int{}
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		jname, _ := findJsonName(tag)
		fieldNames[jname] = i
	}
	fieldNum, ok := fieldNames[fn]
	if !ok {
		return fmt.Errorf("field %s does not exist within the provided item", fn)
	}

	fieldVal := v.Field(fieldNum)
	fieldVal.Set(reflect.ValueOf(fv))
	return nil
}
func HandleJSONObject(object interface{}, key, indentation string) {
	switch t := object.(type) {
	case string:
		fmt.Println(indentation+key+": ", t) // t has type string
	case bool:
		fmt.Println(indentation+key+": ", t) // t has type bool
	case float64:
		fmt.Println(indentation+key+": ", t) // t has type float64
	case map[string]interface{}:
		fmt.Println(indentation + key + ":")
		for k, v := range t {
			HandleJSONObject(v, k, indentation+"\t")
		}
	case []interface{}:
		fmt.Println(indentation + key + ":")
		for index, v := range t {
			HandleJSONObject(v, "["+strconv.Itoa(index)+"]", indentation+"\t")
		}
	}
}
