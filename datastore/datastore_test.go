package datastore

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestInitialize(t *testing.T) {
	ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println(exPath)

	type args struct {
		connection string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Test cases.
		{name: "empty string", args: args{connection: ""}, wantErr: true},
		{name: "invalid path", args: args{connection: "c:/test/fake/path"}, wantErr: true},
		{name: "Valid path", args: args{connection: exPath}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(tt.args.connection); (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
