package rule

import "github.com/aynakeya/deepcolor"

type CommonSearchRules struct {
	SearchID    deepcolor.Item
	SearchTitle deepcolor.Item
	SearchCover deepcolor.Item
	SearchYear  deepcolor.Item
	SearchTag   deepcolor.Item
	SearchDesc  deepcolor.Item
}

type CommonInfoRules struct {
	InfoTitle deepcolor.Item
	InfoCover deepcolor.Item
	InfoYear  deepcolor.Item
	InfoTag   deepcolor.Item
	InfoDesc  deepcolor.Item
}

type CommonPlaylistRules struct {
	PlaylistNames deepcolor.Item
	Videos        deepcolor.Item
	VideoNames    deepcolor.Item
}
