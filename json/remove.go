package json

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strconv"
)

func Remove(pointer string, json []byte) ([]byte, error) {
	p, err := CreatePathPatch("remove", pointer)
	if err != nil {
		return json, err
	}
	return p.ApplyIndent(json, "  ")
}

// RemoveByValue removes a given value if present at the pointer in the given json.
// At this point this method supports removing a number from a slice.
func RemoveByValue(pointer string, value string, json []byte) ([]byte, error) {
	var inputNumValue int
	var inputStrValue string
	if value != "" && value[0] != '"' {
		err := jsoniter.Unmarshal([]byte(value), &inputNumValue)
		if err != nil {
			return json, err
		}
	} else {
		err := jsoniter.Unmarshal([]byte(value), &inputStrValue)
		if err != nil {
			return json, err
		}
	}
	foundValue, kind, err := Value(pointer, string(json))
	if err != nil {
		return json, err
	}
	switch kind {
	case reflect.Slice:
		arr := foundValue.([]interface{})
		for i := 0; i < len(arr); i++ {
			switch v := arr[i].(type) {
			case float64:
				jsonIntValue := int(v)
				if jsonIntValue == inputNumValue {
					return Remove(pointer+"/"+strconv.Itoa(i), json)
				}
			case string:
				jsonStringValue := v
				if jsonStringValue == inputStrValue {
					return Remove(pointer+"/"+strconv.Itoa(i), json)
				}
			default:
				continue
			}
		}
	case reflect.String:
		jsonStrValue := foundValue.(string)
		if jsonStrValue == inputStrValue {
			return Remove(pointer, json)
		}
	case reflect.Float64:
		jsonIntValue := int(foundValue.(float64))
		if jsonIntValue == inputNumValue {
			return Remove(pointer, json)
		}
	}
	return json, nil
}
