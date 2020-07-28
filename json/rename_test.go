package json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRename(t *testing.T) {
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
			name: "community story rename",
			args: args{
				pointer: "/apiName",
				value:   "$id",
				json: []byte(`{
  "$id": "communityStory",
  "$schema": "https://schema.yext.com/config/km/entity-type/v1",
  "apiName": "ce_communityStory",
  "enabled": true,
  "field": [
    {
      "id": "yext/landingPageUrl",
      "required": false
    },
    {
      "id": "yext/organizerEmail",
      "required": false
    },
    {
      "id": "yext/description",
      "required": false
    },
    {
      "id": "yext/name",
      "required": true
    },
    {
      "id": "goal",
      "required": false
    },
    {
      "id": "localNonProfits",
      "required": false
    },
    {
      "id": "primaryCTA",
      "required": false
    },
    {
      "id": "promotingEvents",
      "required": false
    },
    {
      "id": "territory",
      "required": false
    },
    {
      "id": "theme",
      "required": false
    },
    {
      "id": "yext/folder",
      "required": false
    },
    {
      "id": "yext/labels",
      "required": false
    },
    {
      "id": "yext/entityId",
      "required": true
    },
    {
      "id": "yext/videos",
      "required": false
    }
  ],
  "description": "This is to highlight local initiatives Turtlehead Tacos is doing in the neighbourhood.",
  "displayName": "Community Story",
  "pluralDisplayName": "Community Stories"
}`),
			},
			want: []byte(`{
  "$id": "ce_communityStory",
  "$schema": "https://schema.yext.com/config/km/entity-type/v1",
  "enabled": true,
  "field": [
    {
      "id": "yext/landingPageUrl",
      "required": false
    },
    {
      "id": "yext/organizerEmail",
      "required": false
    },
    {
      "id": "yext/description",
      "required": false
    },
    {
      "id": "yext/name",
      "required": true
    },
    {
      "id": "goal",
      "required": false
    },
    {
      "id": "localNonProfits",
      "required": false
    },
    {
      "id": "primaryCTA",
      "required": false
    },
    {
      "id": "promotingEvents",
      "required": false
    },
    {
      "id": "territory",
      "required": false
    },
    {
      "id": "theme",
      "required": false
    },
    {
      "id": "yext/folder",
      "required": false
    },
    {
      "id": "yext/labels",
      "required": false
    },
    {
      "id": "yext/entityId",
      "required": true
    },
    {
      "id": "yext/videos",
      "required": false
    }
  ],
  "description": "This is to highlight local initiatives Turtlehead Tacos is doing in the neighbourhood.",
  "displayName": "Community Story",
  "pluralDisplayName": "Community Stories"
}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Rename(tt.args.pointer, tt.args.value, tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, string(tt.want), string(got))
		})
	}
}
