package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSum(t *testing.T) {
	type args struct {
		filePath string
		option   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "sha256 successfully",
			args: args{
				filePath: "./mockdata/file.txt",
				option:   "sha256",
			},
			want:    "518bf3bb2f528e28df49ae536de2422bcd776f076885057bffefb8ffa4e9340b",
			wantErr: false,
		},
		{
			name: "sha1 successfully",
			args: args{
				filePath: "./mockdata/file.txt",
				option:   "sha1",
			},
			want:    "34ec281c0112f4642ed06725b342caa9879f3596",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckSum(tt.args.filePath, tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *got != tt.want {
				t.Errorf("CheckSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCheckSumCommand(t *testing.T) {
	cmd := NewCheckSumCommand()
	cmd.SetArgs([]string{
		"--file", "./mockdata/file.txt", "md5",
	})
	err := cmd.Execute()
	assert.Nil(t, err)
}
