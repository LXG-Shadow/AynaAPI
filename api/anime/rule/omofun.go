package rule

import "github.com/aynakeya/deepcolor"

type OmofunRules struct {
	CommonSearchRules
	CommonInfoRules
	CommonPlaylistRules
	VideoId     deepcolor.Item
	VideoEncUrl deepcolor.Item
	VideoAesIv  deepcolor.Item
}

func InitializeOmofunRules() OmofunRules {
	return OmofunRules{
		CommonSearchRules: CommonSearchRules{
			SearchID: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.AttributeSelector(
							"div.module-card-item-info > div.module-card-item-title > a",
							"href"),
						Substitution: map[string]string{
							".html":                     "",
							"/index.php/vod/detail/id/": "",
						},
					},
				},
			},
			SearchTitle: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-card-item-info > div.module-card-item-title > a > strong"),
					},
				},
			},
			SearchCover: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.AttributeSelector(
							"div.module-item-pic > img",
							"data-original"),
					},
				},
			},
			SearchYear: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-card-item-info > div:nth-child(2) > div",
						),
						Substitution: map[string]string{
							" ":       "",
							"[^1-9]*": "",
						},
					},
				},
			},
			SearchTag: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-card-item-class"),
					},
				},
			},
			SearchDesc: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-card-item-info > div:nth-child(3) > div"),
					},
				},
			},
		},
		CommonInfoRules: CommonInfoRules{
			InfoTitle: deepcolor.Item{
				Type: deepcolor.ItemTypeSingle,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector("div.module-main > div.module-info-main > div.module-info-heading > h1"),
					},
				},
			},
			InfoCover: deepcolor.Item{
				Type: deepcolor.ItemTypeSingle,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.AttributeSelector(
							"div.module-main > div.module-info-poster > div > div > img",
							"data-original"),
					},
				},
			},
			InfoYear: deepcolor.Item{
				Type: deepcolor.ItemTypeSingle,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-main > div.module-info-main > div.module-info-heading > div.module-info-tag > div:nth-child(1)"),
					},
				},
			},
			InfoTag: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-main > div.module-info-main > div.module-info-heading > div.module-info-tag > div:nth-child(3) > a"),
					},
				},
			},
			InfoDesc: deepcolor.Item{
				Type: deepcolor.ItemTypeSingle,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-main > div.module-info-main > div.module-info-content > div.module-info-items > div.module-info-item.module-info-introduction > div > p"),
					},
				},
			},
		},
		CommonPlaylistRules: CommonPlaylistRules{
			PlaylistNames: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-tab-items-box > div.module-tab-item > span"),
					},
				},
			},
			VideoNames: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"a.module-play-list-link > span"),
					},
				},
			},
			Videos: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.AttributeSelector(
							"a.module-play-list-link",
							"href"),
						Substitution: map[string]string{
							"/index.php/vod/play/id/": "",
							".html":                   "",
							"/nid/":                   "-",
							"/sid/":                   "-",
						},
					},
				},
			},
		},
		VideoId: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.RegExpSelector(
						"\"url\":\"([^\"]*)\",\"url_next\""),
					Substitution: map[string]string{
						"\"url\":\"":      "",
						"\",\"url_next\"": "",
					},
				},
			},
		},
		VideoAesIv: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.RegExpSelector(
						"bt_token = \"([^;]*)\";"),
					Substitution: map[string]string{
						"bt_token = \"": "",
						"\";":           "",
					},
				},
			},
		},
		VideoEncUrl: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.RegExpSelector(
						"getVideoInfo\\(\"[^\"]*\"\\)"),
					Substitution: map[string]string{
						"getVideoInfo\\(\"": "",
						"\"\\)":             "",
					},
				},
			},
		},
	}
}
