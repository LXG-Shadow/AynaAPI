package rule

import "github.com/aynakeya/deepcolor"

type DldmRules struct {
	CommonSearchRules
	CommonInfoRules
	CommonPlaylistRules
	VideoId  deepcolor.Item
	VideoUrl deepcolor.Item
}

func InitializeDldmRules() DldmRules {
	return DldmRules{
		CommonSearchRules: CommonSearchRules{
			SearchID: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.AttributeSelector(
							"div.module-search-item > div.video-info > div.video-info-header > h3 > a",
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
							"div.module-search-item > div.video-info > div.video-info-header > h3 > a"),
					},
				},
			},
			SearchCover: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.AttributeSelector(
							"div.module-search-item > div.video-cover > div > div > img",
							"data-src"),
					},
				},
			},
			SearchYear: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-search-item > div.video-info > div.video-info-header > div > div:nth-child(2)",
						),
						Substitution: map[string]string{
							" ": "",
						},
					},
				},
			},
			SearchTag: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-search-item >  div.video-info > div.video-info-header > div > a > span"),
					},
				},
			},
			SearchDesc: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-search-item >  div.video-info > div.video-info-main > div:nth-child(3) > div"),
					},
				},
			},
		},
		CommonInfoRules: CommonInfoRules{
			InfoTitle: deepcolor.Item{
				Type: deepcolor.ItemTypeSingle,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(".page-title"),
					},
				},
			},
			InfoCover: deepcolor.Item{
				Type: deepcolor.ItemTypeSingle,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.AttributeSelector(
							"#main > div > div.box.view-heading > div.video-cover > div > div > img",
							"data-src"),
					},
				},
			},
			InfoYear: deepcolor.Item{
				Type: deepcolor.ItemTypeSingle,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"#main > div > div.box.view-heading > div.video-info > div.video-info-header > div.scroll-box > div > a:nth-child(3)"),
					},
				},
			},
			InfoTag: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"#main > div > div.box.view-heading > div.video-info > div.video-info-header > div.scroll-box > div > div > a"),
					},
				},
			},
			InfoDesc: deepcolor.Item{
				Type: deepcolor.ItemTypeSingle,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector("#main > div > div.box.view-heading > div.video-info > div.video-info-main > div.video-info-items > div > span"),
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
							"div.module-heading > div.module-tab.module-player-tab > div > div.module-tab-content > div > span"),
					},
				},
			},
			VideoNames: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.TextSelector(
							"div.module-blocklist.scroll-box.scroll-box-y > div > a > span"),
					},
				},
			},
			Videos: deepcolor.Item{
				Type: deepcolor.ItemTypeList,
				Rules: []deepcolor.ItemRule{
					{
						Selector: deepcolor.AttributeSelector(
							"div.module-blocklist.scroll-box.scroll-box-y > div > a",
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
		VideoUrl: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.JsonSelector("url"),
					Substitution: map[string]string{
						"\\/": "/",
					},
				},
			},
		},
	}
}
