package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
