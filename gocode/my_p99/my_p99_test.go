package my_p99

import (
	"strconv"
	"testing"
)

func Test_getIndex(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		struct {
			name string
			args args
			want int
		}{name: "val0", args: struct{ val int64 }{val: 0}, want: 0},
		{name: "val1", args: struct{ val int64 }{val: 1}, want: 0},
		{name: "val10", args: struct{ val int64 }{val: 10}, want: 1},
		{name: "val900", args: struct{ val int64 }{val: 900}, want: 90},
		{name: "val909", args: struct{ val int64 }{val: 909}, want: 90},
		{name: "val999", args: struct{ val int64 }{val: 999}, want: 99},
		{name: "two1000", args: struct{ val int64 }{val: 1000}, want: 100},
	}
	for i := 0; i < 2000; i += 8 {
		tests = append(tests, struct {
			name string
			args args
			want int
		}{name: "test" + strconv.Itoa(i), args: struct{ val int64 }{val: int64(i)}, want: -1})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIndex(tt.args.val); got != tt.want {
				t.Errorf("getIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
