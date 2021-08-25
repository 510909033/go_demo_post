package my_benchmark

import "testing"

/*
测试时间默认是1秒

 go test -bench=. -run=none
 go test -test.v -test.bench ^BenchmarkAdd$ -test.run ^$

*/

func TestAdd(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{struct {
		name string
		args args
		want int
	}{name: "a1", args: args{[]int{1, 2, 3, 4}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a...); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2, 3, 4)
	}
}
