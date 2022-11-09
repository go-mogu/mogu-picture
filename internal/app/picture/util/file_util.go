package util

import "github.com/gogf/gf/v2/os/gfile"

func GetPicExpandedName(oldName string) string {
	picExpandedName := gfile.ExtName(oldName)
	if picExpandedName == "" {
		picExpandedName = "jpg"
	}
	return picExpandedName
}
