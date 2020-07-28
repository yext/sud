package json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	type args struct {
		pointer string
		value   string
		json    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "add product features to dependencies.json",
			args: args{
				pointer: "/productFeatureIds/-",
				value:   "304",
				json: []byte(`{
  "productFeatureIds": [
    1,
    2
  ]
}`),
			},
			want: []byte(`{
  "productFeatureIds": [
    1,
    2,
    304
  ]
}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.pointer, tt.args.value, tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, string(tt.want), string(got))
		})
	}
}
