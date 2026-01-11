package cmd

import (
	"bytes"
	"testing"
)

func TestCd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "valid",
			args:    []string{"cd", "my-workspace"},
			wantErr: false,
		},
		{
			name:    "valid with flag",
			args:    []string{"cd", "my-workspace", "--new"},
			wantErr: false,
		},
		{
			name:    "missing arg",
			args:    []string{"cd"},
			wantErr: true,
		},
		{
			name:    "too many args",
			args:    []string{"cd", "ws1", "ws2"},
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