package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func ToMd5Str(val []byte) string {
	md := md5.New()
	md.Write(val)
	return hex.EncodeToString(md.Sum(nil))
}

func FileSizeFormat(size int64) string {
	if size/1024 < 1024 {
		return fmt.Sprintf("%dK", size/1024)
	}
	return fmt.Sprintf("%dM", size/(1024*1024))
}
