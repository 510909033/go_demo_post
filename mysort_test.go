package main

import (
	"reflect"
	"testing"
)

func Test_maopao1(t *testing.T) {
	type args struct {
		arr []int
	}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		struct {
			name string
			args args
		}{name: "test nil", args: args{arr: []int{1, 3, 5, 7, 9, 8, 6, 4, 2, 0}}},
	}

	t.Log("want:", want)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var before = make([]int, len(tt.args.arr), 100)
			copy(before, tt.args.arr)
			maopao1(tt.args.arr)
			t.Logf("before:%+v", before)
			t.Log("result:", tt.args.arr)
			if !reflect.DeepEqual(want, tt.args.arr) {
				t.Error("和want不一致")
			}
		})
	}

}
