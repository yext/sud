package json

func Move(pointer string, toPtr string, json []byte) ([]byte, error) {
	p, err := CreatePatch("move", pointer, toPtr)
	if err != nil {
		return json, err
	}
	return p.ApplyIndent(json, "  ")
}
