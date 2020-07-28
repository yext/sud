package json

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValue(t *testing.T) {
	type args struct {
		pointer  string
		jsonText string
	}
	tests := []struct {
		name      string
		args      args
		wantValue interface{}
		wantKind  reflect.Kind
		wantErr   bool
	}{
		{
			name: "get value of an array",
			args: args{
				pointer: "/productFeatureIds",
				jsonText: `{
  "productFeatureIds": [100, 200]
}`,
			},
			wantValue: []interface{}{100.0, 200.0},
			wantKind:  reflect.Slice,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotKind, err := Value(tt.args.pointer, tt.args.jsonText)
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.wantValue, gotValue)
			require.Equal(t, tt.wantKind, gotKind)
		})
	}
}
