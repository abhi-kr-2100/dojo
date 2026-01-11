package cmd

import (
	"bytes"
	"testing"
)

func TestRm(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "valid with arg",
			args:    []string{"rm", "my-workspace"},
			wantErr: false,
		},
		{
			name:    "valid without arg",
			args:    []string{"rm"},
			wantErr: false,
		},
		{
			name:    "too many args",
			args:    []string{"rm", "ws1", "ws2"},
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