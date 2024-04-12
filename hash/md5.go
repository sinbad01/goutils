package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileMD5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New(
			fmt.Sprintf("md5.go hash.FileMD5 os open error %v", err),
		)
	}

	md5 := md5.New()
	_, err = io.Copy(md5, file)
	if err != nil {
		return "", errors.New(
			fmt.Sprintf("md5.go hash.FileMD5 os copy error %v", err),
		)
	}

	return hex.EncodeToString(md5.Sum(nil)), nil
}

func StringMD5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	return hex.EncodeToString(md5.Sum(nil))
}
