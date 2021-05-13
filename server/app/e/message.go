package e

var CODE_MESSAGE = map[int]string{
	API_ERROR_UNKNOWN:     "Unknown error",
	API_OK:                "Get success",
	API_REQUIRE_PARAMETER: "Require parameter",
	API_INVALID_PARAMETER: "Invalid parameter",

	BGM_INITIALIZE_FAIL:        "获取信息失败",
	BGM_PROVIDER_NOT_AVAILABLE: "该来源暂时不可用",
	BGM_SEARCH_FAIL:            "资源搜索失败",
	BGM_NO_RESULT:              "无结果",
}

// GetMessage get error information based on Code
func GetMessage(code int) string {
	msg, ok := CODE_MESSAGE[code]
	if ok {
		return msg
	}
	return CODE_MESSAGE[API_ERROR_UNKNOWN]
}
