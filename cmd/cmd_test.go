package cmd

import (
	"testing"
)

func TestCommandsExist(t *testing.T) {
	rootCmd := NewRootCmd()
	
	tests := []struct {
		name string
	}{
		{"ls"},
		{"new"},
		{"cd"},
		{"rm"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := rootCmd.Find([]string{tt.name})
			if err != nil {
				t.Errorf("Command %q not found: %v", tt.name, err)
			}
		})
	}
}
