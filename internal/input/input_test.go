package input

import (
	"reflect"
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func TestFiles(t *testing.T) {
	type args struct {
		srcPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Path
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"a",
			args{"./"},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Files(tt.args.srcPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Files() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Files() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dirFiles(t *testing.T) {
	type args struct {
		srcPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Path
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dirFiles(tt.args.srcPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("dirFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dirFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_file(t *testing.T) {
	type args struct {
		srcPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Path
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := file(tt.args.srcPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("file() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("file() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHiddenFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHiddenFile(tt.args.path); got != tt.want {
				t.Errorf("isHiddenFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
