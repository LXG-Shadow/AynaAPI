package resp

import "AynaAPI/server/app"

type PublicInfo struct {
	Username        string `json:"username"`
	PermissionLevel int    `json:"permission"`
}

type UserPublicInfo struct {
	app.AppJsonResponse
	Data PublicInfo
}

type LoginResp struct {
	app.AppJsonResponse
	Data map[string]string `example:"token:token1234"`
}
