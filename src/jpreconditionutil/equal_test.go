package jprecondition

import "testing"

func TestIsEqual(t *testing.T) {
	type args struct {
		flag bool
		msg  interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
