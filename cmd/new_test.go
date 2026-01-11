package cmd

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "valid",
			args:    []string{"new", "my-workspace"},
			wantErr: false,
		},
		{
			name:    "missing arg",
			args:    []string{"new"},
			wantErr: true,
		},
		{
			name:    "too many args",
			args:    []string{"new", "ws1", "ws2"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd := NewRootCmd()
			rootCmd.SetOut(buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}