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
		return fmt.Sprintf("%.2fK", float64(size)/1024)
	}
	return fmt.Sprintf("%.2fM", float64(size)/(1024*1024))
}
