package json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplace(t *testing.T) {
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
			name: "replaces value",
			args: args{
				pointer: "/PageSetConfig/ContentMapping/Event",
				value:   "\"VirtualOnward\"",
				json: []byte(`{
  "PageSetConfig": {
    "ContentMapping": {
	  "Event": "Onward"
	}
  }
}`),
			},
			want: []byte(`{
  "PageSetConfig": {
    "ContentMapping": {
      "Event": "VirtualOnward"
    }
  }
}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Replace(tt.args.pointer, tt.args.value, tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("Replace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, string(tt.want), string(got))
		})
	}
}
