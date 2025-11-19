package lib

import (
	"fmt"
	"reflect"
)

func RemoveNulls(m map[string]interface{}) (data map[string]interface{}) {
	val := reflect.ValueOf(m)
	for _, e := range val.MapKeys() {
		v := val.MapIndex(e)
		if v.IsNil() {
			delete(m, e.String())
			continue
		}
		switch t := v.Interface().(type) {
		// If key is a JSON object (Go Map), use recursion to go deeper
		case map[string]interface{}:
			RemoveNulls(t)
		}
	}
	return data
}

func RemoveNull(maps interface{}, field string) (result interface{}) {
	newData := map[string]interface{}{
		"KOSTLGS": "NULL",
	}

	data := maps.(map[string]interface{})[field]
	val := reflect.TypeOf(data)
	switch val {
	case nil:
		fmt.Println("nil")
		return newData
	default:
		fmt.Println("default")
		return data
	}
}

// case val.Kind() == reflect.Int:
// 	fmt.Println("int")
// 	return data
// case val.Kind() == reflect.String:
// 	fmt.Println("string")
// 	return data
// case val.Kind() == reflect.Bool:
// 	fmt.Println("bool")
// 	return data
// case val.Kind() == reflect.Interface:
// 	fmt.Println("interface")
// 	return data
// case val.Kind() == reflect.Map:
// 	return data
// 	fmt.Println("map")
