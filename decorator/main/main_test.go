package main

import "testing"

func TestDemoDecorator(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		struct{ name string }{name: "main_test_demo_decorator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DemoDecorator()
		})
	}
}
