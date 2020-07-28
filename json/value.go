package json

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/xeipuuv/gojsonpointer"
)

func Value(pointer string, jsonText string) (interface{}, reflect.Kind, error) {
	var jsonDocument map[string]interface{}
	err := jsoniter.Unmarshal([]byte(jsonText), &jsonDocument)
	if err != nil {
		return nil, reflect.Invalid, err
	}

	p, err := gojsonpointer.NewJsonPointer(pointer)
	if err != nil {
		return nil, reflect.Invalid, err
	}

	v, kind, err := p.Get(jsonDocument)
	if err != nil {
		return nil, reflect.Invalid, err
	}
	return v, kind, nil
}
