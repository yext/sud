package json

func Replace(pointer string, value string, json []byte) ([]byte, error) {
	p, err := CreateValuePatch("replace", pointer, value)
	if err != nil {
		return json, err
	}
	return p.ApplyIndent(json, "  ")
}
