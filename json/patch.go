package json

import (
	"github.com/saurabhaditya/json-patch"
)

func CreatePatch(operation string, from string, path string) (jsonpatch.Patch, error) {
	return jsonpatch.DecodePatch([]byte(`[
 {"op": "` + operation + `", "from": "` + from + `", "path": "` + path + `"}
]`))
}

func CreateValuePatch(op string, path string, value string) (jsonpatch.Patch, error) {
	return jsonpatch.DecodePatch([]byte(`[
 {"op": "` + op + `", "path": "` + path + `", "value": ` + value + `}
]`))
}

func CreatePathPatch(op string, path string) (jsonpatch.Patch, error) {
	return jsonpatch.DecodePatch([]byte(`[
  {"op": "` + op + `", "path": "` + path + `"}
]`))
}
