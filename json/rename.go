package json

import "path/filepath"

func Rename(pointer string, value string, json []byte) ([]byte, error) {
	toPtr := filepath.Join(filepath.Dir(pointer), value)
	p, err := CreatePatch("move", pointer, toPtr)
	if err != nil {
		return json, err
	}
	return p.ApplyIndent(json, "  ")
}
