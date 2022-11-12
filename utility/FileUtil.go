package utils

import (
	"github.com/go-mogu/mogu-picture/internal/consts/Constants"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	IMAGE_TYPE = 1
	DOC_TYPE   = 2
	VIDEO_TYPE = 3
	MUSIC_TYPE = 4
	OTHER_TYPE = 5
)

var (
	IMG_FILE = []string{Constants.FILE_SUFFIX_BMP, Constants.FILE_SUFFIX_JPG,
		Constants.FILE_SUFFIX_PNG,
		Constants.FILE_SUFFIX_TIF,
		Constants.FILE_SUFFIX_GIF,
		Constants.FILE_SUFFIX_JPEG,
		Constants.FILE_SUFFIX_WEBP}
	DOC_FILE = []string{
		Constants.FILE_SUFFIX_DOC,
		Constants.FILE_SUFFIX_DOCX,
		Constants.FILE_SUFFIX_TXT,
		Constants.FILE_SUFFIX_HLP,
		Constants.FILE_SUFFIX_WPS,
		Constants.FILE_SUFFIX_RTF,
		Constants.FILE_SUFFIX_XLS,
		Constants.FILE_SUFFIX_XLSX,
		Constants.FILE_SUFFIX_PPT,
		Constants.FILE_SUFFIX_PPTX,
		Constants.FILE_SUFFIX_JAVA,
		Constants.FILE_SUFFIX_HTML,
		Constants.FILE_SUFFIX_PDF,
		Constants.FILE_SUFFIX_MD,
		Constants.FILE_SUFFIX_SQL,
		Constants.FILE_SUFFIX_CSS,
		Constants.FILE_SUFFIX_JS,
		Constants.FILE_SUFFIX_VUE,
		Constants.FILE_SUFFIX_JAVA}
	VIDEO_FILE = []string{
		Constants.FILE_SUFFIX_AVI,
		Constants.FILE_SUFFIX_MP4,
		Constants.FILE_SUFFIX_MPG,
		Constants.FILE_SUFFIX_MOV,
		Constants.FILE_SUFFIX_SWF,
		Constants.FILE_SUFFIX_3GP,
		Constants.FILE_SUFFIX_RM,
		Constants.FILE_SUFFIX_RMVB,
		Constants.FILE_SUFFIX_WMV,
		Constants.FILE_SUFFIX_MKV}
	MUSIC_FILE = []string{
		Constants.FILE_SUFFIX_WAV,
		Constants.FILE_SUFFIX_AIF,
		Constants.FILE_SUFFIX_AU,
		Constants.FILE_SUFFIX_MP3,
		Constants.FILE_SUFFIX_RAM,
		Constants.FILE_SUFFIX_WMA,
		Constants.FILE_SUFFIX_MMF,
		Constants.FILE_SUFFIX_AMR,
		Constants.FILE_SUFFIX_AAC,
		Constants.FILE_SUFFIX_FLAC}
	ALL_FILE = []string{}
)

func init() {
	ALL_FILE = append(ALL_FILE, IMG_FILE...)
	ALL_FILE = append(ALL_FILE, DOC_FILE...)
	ALL_FILE = append(ALL_FILE, VIDEO_FILE...)
	ALL_FILE = append(ALL_FILE, MUSIC_FILE...)
}

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

func GetPicExpandedName(oldName string) string {
	picExpandedName := gfile.ExtName(oldName)
	if picExpandedName == "" {
		picExpandedName = "jpg"
	}
	return picExpandedName
}

// IsSafe 通过其后缀名判断其是否合法,合法后缀名为常见的
// suffix 后缀名
// 合法返回true，不合法返回false
func IsSafe(suffix string) bool {
	suffix = gstr.ToLower(suffix)
	for _, s := range ALL_FILE {
		if s == suffix {
			return true
		}
	}
	return false

}

func GetFileExtendsByType(fileType int) (fileExtends []string) {
	fileExtends = make([]string, 0)
	switch fileType {
	case IMAGE_TYPE:
		fileExtends = IMG_FILE
	case DOC_TYPE:
		fileExtends = DOC_FILE
	case VIDEO_TYPE:
		fileExtends = VIDEO_FILE
	case MUSIC_TYPE:
		fileExtends = MUSIC_FILE
	case OTHER_TYPE:
		fileExtends = ALL_FILE
	}
	return fileExtends
}

func PathSplitFormat(filePath string) string {
	filePath = gstr.Replace(filePath, "///", "/")
	filePath = gstr.Replace(filePath, "//", "/")
	filePath = gstr.Replace(filePath, "\\\\\\", "\\")
	filePath = gstr.Replace(filePath, "\\\\", "\\")
	return filePath
}
