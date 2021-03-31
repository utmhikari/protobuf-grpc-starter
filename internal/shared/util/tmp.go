package util

import (
	"errors"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/util/fs"
	"os"
)


func GetTmpDir() string {
	return "./tmp"
}


func MakeTmpDir() error {
	p := GetTmpDir()
	if fs.IsFile(p) {
		return errors.New("tmp is a file")
	}
	if fs.IsDirectory(p) {
		return nil
	}

	err := os.MkdirAll(p, 0777)
	if err != nil {
		return err
	}

	return nil
}
