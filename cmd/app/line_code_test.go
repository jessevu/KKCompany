package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineCount(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "successfully",
			args: args{
				filePath: "./mockdata/file.txt",
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LineCount(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *got != tt.want {
				t.Errorf("LineCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLineCountCommand(t *testing.T) {
	cmd := NewLineCountCommand()
	cmd.SetArgs([]string{
		"--file", "./mockdata/file.txt",
	})
	err := cmd.Execute()
	assert.Nil(t, err)
}
