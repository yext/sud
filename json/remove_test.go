package json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemove(t *testing.T) {
	type args struct {
		pointer string
		json    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "remove product feature from dependencies.json",
			args: args{
				pointer: "/productFeatureIds/2",
				json: []byte(`{
  "productFeatureIds": [
    1,
    2,
	3
  ]
}`),
			},
			want: []byte(`{
  "productFeatureIds": [
    1,
    2
  ]
}`),
			wantErr: false,
		},
		{
			name: "remove apiName",
			args: args{
				pointer: "/apiName",
				json: []byte(`{
  "$id": "activeOnAnswers",
  "$schema": "https://schema.yext.com/config/km/field/v1",
  "apiName": "c_activeOnAnswers",
  "description": "",
  "displayName": "Active on Answers"
}`),
			},
			want: []byte(`{
  "$id": "activeOnAnswers",
  "$schema": "https://schema.yext.com/config/km/field/v1",
  "description": "",
  "displayName": "Active on Answers"
}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Remove(tt.args.pointer, tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, string(tt.want), string(got))
		})
	}
}

func TestRemoveByValue(t *testing.T) {
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
			name: "remove product feature by value from dependencies.json",
			args: args{
				pointer: "/productFeatureIds",
				value:   "2",
				json: []byte(`{
  "productFeatureIds": [
    1,
    2,
    3
  ]
}`),
			},
			want: []byte(`{
  "productFeatureIds": [
    1,
    3
  ]
}`),
			wantErr: false,
		},
		{
			name: "remove /type/desc if matches",
			args: args{
				pointer: "/type/desc",
				value:   "\"some type description\"",
				json: []byte(`{
  "$id": "123",
  "type": {
    "desc": "some type description"
  }
}`),
			},
			want: []byte(`{
  "$id": "123",
  "type": {}
}`),
			wantErr: false,
		},
		{
			name: "do not remove /type/desc if does not match",
			args: args{
				pointer: "/type/desc",
				value:   "\"some type description\"",
				json: []byte(`{
  "$id": "123",
  "type": {
    "desc": "some other description"
  }
}`),
			},
			want: []byte(`{
  "$id": "123",
  "type": {
    "desc": "some other description"
  }
}`),
			wantErr: false,
		},
		{
			name: "remove int value if matches",
			args: args{
				pointer: "/rank",
				value:   "5",
				json: []byte(`{
  "rank": 5
}`),
			},
			want:    []byte(`{}`),
			wantErr: false,
		},
		{
			name: "does not remove int value if no match",
			args: args{
				pointer: "/rank",
				value:   "4",
				json: []byte(`{
  "rank": 5
}`),
			},
			want: []byte(`{
  "rank": 5
}`),
			wantErr: false,
		},
		{
			name: "remove string by value from array",
			args: args{
				pointer: "/tools",
				value:   "\"sud\"",
				json: []byte(`{
  "tools": ["sud", "yext"]
}`),
			},
			want: []byte(`{
  "tools": [
    "yext"
  ]
}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveByValue(tt.args.pointer, tt.args.value, tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveByValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
