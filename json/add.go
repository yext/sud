package json

func Add(pointer string, value string, json []byte) ([]byte, error) {
	p, err := CreateValuePatch("add", pointer, value)
	if err != nil {
		return json, err
	}
	return p.ApplyIndent(json, "  ")
}
