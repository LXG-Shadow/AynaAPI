package vhttp

import (
	"AynaAPI/utils/vstring"
	"net/url"
	"path"
	"regexp"
)

func IsUrl(url string) bool {
	urlRegExp := regexp.MustCompile(
		"(?i)^(?:http|ftp)s?://" +
			"(?:(?:[A-Z0-9](?:[A-Z0-9-]{0,61}[A-Z0-9])?\\.)+(?:[A-Z]{2,6}\\.?|[A-Z0-9-]{2,}\\.?)|" +
			"localhost|" +
			"\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3})" +
			"(?::\\d+)?" +
			"(?:/?|[/?]\\S+)$")
	return urlRegExp.FindString(url) != ""
}

func QueryEscapeWithEncoding(str string, encoding string) string {
	return url.QueryEscape(EncodeString(str, encoding))
}

func QueryUnescapeWithEncoding(str string, encoding string) string {
	s, err := url.QueryUnescape(EncodeString(str, encoding))
	if err != nil {
		return ""
	}
	return s
}

func GetUrlHost(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	return u.Host
}

func GetUrlPath(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	return u.Path
}

func JoinUrl(base string, paths ...string) string {
	u, err := url.Parse(base)
	if err != nil {
		return base
	}
	curPath := u.Path
	flag := false
	if len(paths) > 0 {
		if last := paths[len(paths)-1]; last[len(last)-1:] == "/" {
			flag = true
		}
	}
	newPath := path.Join(paths...)
	if vstring.LenString(newPath) < 0 {
		return u.String()
	}
	if newPath[:1] == "/" {
		curPath = "/"
	}
	u.Path = path.Join(curPath, newPath)
	if flag {
		return u.String() + "/"
	} else {
		return u.String()
	}
}

func CompleteUrl(host, path string) string {
	if path == "" {
		return ""
	}
	if !IsUrl(path) {
		return JoinUrl(host, path)
	}
	return path
}
