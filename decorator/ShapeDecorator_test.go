package decorator

import "testing"

func TestDemoDecorator(t *testing.T) {
	tests := []struct {
		name string
	}{
		struct{ name string }{name: "test_demo_decorator"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DebugDemoDecorator()
		})
	}
}
