package split

import "strings"

// Split 将s按照sep进行分割，返回一个字符串切片
// Split（"我爱你","爱"） => ["我", "你"]
func Split(s string, sep string) (ret []string) {
	ret = make([]string, 0, strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
