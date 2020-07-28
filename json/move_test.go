package json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMove(t *testing.T) {
	type args struct {
		pointer string
		toPtr   string
		json    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "rename name to $id",
			args: args{
				pointer: "/apiName",
				toPtr:   "/$id",
				json: []byte(`{
  "$id": "price",
  "apiName": "c_price"
}`),
			},
			want: []byte(`{
  "$id": "c_price"
}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Move(tt.args.pointer, tt.args.toPtr, tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, string(tt.want), string(got))
		})
	}
}
