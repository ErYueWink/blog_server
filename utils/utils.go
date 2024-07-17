package utils

// In_list 判断图片后缀名是否合法
func In_list(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}
