package chain

import "testing"

func TestDemoChain(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		struct{ name string }{name: "test_demo_chain"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DemoChain()
		})
	}
}
