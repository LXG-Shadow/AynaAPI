package e

var CODE_MESSAGE = map[int]string{
	SUCCESS:            "Success",
	INTERNAL_ERROR:     "Server internal error",
	UKNOWN_ERROR:       "Unkwown error",
	ERROR_READFILE:     "读取文件出错",
	EXTERNAL_API_ERROR: "第三方api错误",

	NOVEL_PROVIDER_URL_NOT_SUPPORT:   "非书籍源的url",
	NOVEL_PROVIDER_URL_NOT_SUPPORT_2: "url不支持",
}

// GetMessage get error information based on Code
func GetMessage(code int) string {
	msg, ok := CODE_MESSAGE[code]
	if ok {
		return msg
	}
	return CODE_MESSAGE[INTERNAL_ERROR]
}
