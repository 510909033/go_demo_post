package go_mongodb

import "testing"

func TestDemoMyMongodb(t *testing.T) {
	tests := []struct {
		name string
	}{
		struct{ name string }{name: "name-TestDemoMyMongodb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DemoMyMongodb()
		})
	}
}
