package e

var CODE_MESSAGE = map[int]string{
	ERROR_UNKNOWN: "Unknown error",

	API_ERROR_UNKNOWN:                      "Unknown error",
	API_OK:                                 "Get success",
	API_ERROR_REQUIRE_PARAMETER:            "Require parameter",
	API_ERROR_INVALID_PARAMETER:            "Invalid parameter",
	API_ERROR_REQUIRE_TOKEN:                "Require token",
	API_ERROR_INVALID_TOKEN:                "Invalid token",
	API_ERROR_UPLOAD_SAVE_IMAGE_FAIL:       "Fail to upload image",
	API_ERROR_UPLOAD_IMAGE_NOT_FOUND:       "Image not found",
	API_ERROR_UPLOAD_IVALID_IMAGE:          "Not a valid image",
	API_ERROR_UPLOAD_EXCEED_IMAGE_MAX_SIZE: "Exceed image maxsize",

	API_ERROR_PERMISSION_NOT_ALLOWED: "Do not have enough permission",

	BGM_INITIALIZE_FAIL:        "获取信息失败",
	BGM_PROVIDER_NOT_AVAILABLE: "该来源暂时不可用",
	BGM_SEARCH_FAIL:            "资源搜索失败",
	BGM_NO_RESULT:              "无结果",
	BGM_VIDEO_NOT_FOUND:        "未找到id对应的视频",

	AUTH_OK:                     "登录成功",
	AUTH_ERROR_REQUIRE_USERNAME: "需要用户名",
	AUTH_ERROR_REQUIRE_PASSWORD: "需要密码",
	AUTH_ERROR_U_P_NOT_MATCH:    "用户名密码不匹配",

	NOVEL_PROVIDER_NOT_AVAILABLE: "小说源暂时不可用",
	NOVEL_URL_NOT_SUPPORT:        "Url不匹配",
	NOVEL_GET_DATA_FAIL:          "获取数据失败",
	NOVEL_SEARCH_FAIL:            "小说搜索失败",
	NOVEL_INITIALIZE_FAIL:        "初始化数据失败",
	NOVEL_CHAPTER_NOT_FOUND:      "未找到该章节",

	MUSIC_PROVIDER_NOT_AVAILABLE: "音乐源暂时不可用",
	MUSIC_SEARCH_FAIL:            "音乐搜索失败",
	MUSIC_INITIALIZE_FAIL:        "音乐获取失败",
	MUSIC_GET_DATA_FAIL:          "获取音乐数据失败",
}

// GetMessage get error information based on Code
func GetMessage(code int) string {
	msg, ok := CODE_MESSAGE[code]
	if ok {
		return msg
	}
	return CODE_MESSAGE[API_ERROR_UNKNOWN]
}
