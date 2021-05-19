package strategy

import (
	"go_demo_post/design/interf"
	"go_demo_post/design/strategy/cache_client"
	"testing"
)


func TestCacheClient_Get(t *testing.T) {
	type fields struct {
		client interf.ICacheClient
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		struct {
			name   string
			fields fields
			args   args
		}{name: "get", fields: struct{ client interf.ICacheClient }{client: &cache_client.CacheFile{}} , args: args{"haha"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheClient{
				client: tt.fields.client,
			}
			c.Get(tt.args.key)
		})
	}
}

func TestCacheClient_Set(t *testing.T) {
	type fields struct {
		client interf.ICacheClient
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheClient{
				client: tt.fields.client,
			}
			c.Set(tt.args.key,tt.args.val)
		})
	}
}