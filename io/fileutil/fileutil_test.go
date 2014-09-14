package fileutil

import (
	"testing"
)

func TestExist(t *testing.T) {
	path := "/etc/rc.local"
	if !Exists(path) {
		t.Errorf("%s must exists", path)
	}

	if !IsFile(path) {
		t.Errorf("%s must be file", path)
	}

	if IsDir(path) {
		t.Errorf("%s is not dir", path)
	}
}
