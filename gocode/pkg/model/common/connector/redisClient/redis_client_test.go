package redisClient

import (
	"github.com/go-redis/redis"
	"reflect"
	"testing"
	"time"
)

func TestGetRedisClient(t *testing.T) {
	tests := []struct {
		name string
		want *redis.ClusterClient
	}{
		// TODO: Add test cases.
		struct {
			name string
			want *redis.ClusterClient
		}{name: "name_TestGetRedisClient", want: GetRedisClient()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRedisClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetNx(t *testing.T) {
	type args struct {
		key        string
		value      interface{}
		expiration time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "name_TestSetNx", args: args{
			key:        "test_nx",
			value:      100,
			expiration: time.Second,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetNx(tt.args.key, tt.args.value, tt.args.expiration)
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		key        string
		value      interface{}
		expiration time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "name_TestSetNx", args: args{
			key:        "test_nx",
			value:      100,
			expiration: time.Second,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.key, tt.args.value, tt.args.expiration)
		})
	}
}
