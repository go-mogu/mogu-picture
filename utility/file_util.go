package utils

import (
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	IMG_FILE = []string{Constants.FILE_SUFFIX_BMP, Constants.FILE_SUFFIX_JPG,
		Constants.FILE_SUFFIX_PNG,
		Constants.FILE_SUFFIX_TIF,
		Constants.FILE_SUFFIX_GIF,
		Constants.FILE_SUFFIX_JPEG,
		Constants.FILE_SUFFIX_WEBP}
)

// IsPic suffix 后缀名
// 通过其后缀名判断其是否是图片
func IsPic(suffix string) bool {
	suffix = gstr.ToLower(suffix)
	for i := range IMG_FILE {
		if IMG_FILE[i] == suffix {
			return true
		}
	}
	return false

}
