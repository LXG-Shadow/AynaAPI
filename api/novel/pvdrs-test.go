package novel

import (
	"AynaAPI/pkg/deepcolor"
)

var (
	BiqugeARule = NovelProviderRule{
		Title: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#div1 > h1",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Author: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#info > p:nth-of-type(1)",
					Target:   deepcolor.TextTarget(),
					Filters:  []string{"作(.*)者："},
				},
			},
		},
		Cover: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#fmimg > img",
					Target:   deepcolor.AttributeTarget("src"),
				},
			},
		},
		Abstraction: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#intro",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Chapters: deepcolor.Item{
			Type: deepcolor.ItemTypeMapList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#list > dl > dd > a",
					Key:      "name",
					Target:   deepcolor.TextTarget(),
				},
				{
					Selector: "#list > dl > dd > a",
					Key:      "url",
					Target:   deepcolor.AttributeTarget("href"),
				},
			},
		},
		Content: deepcolor.Item{
			Type: deepcolor.ItemTypeMap,
			Rules: []deepcolor.ItemRule{
				{
					Key:      "content",
					Selector: "#content",
					Target:   deepcolor.TextTarget(),
					Filters:  []string{"\u00a0"},
				},
				{
					Key:      "name",
					Selector: ".bookname > h1",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Search: deepcolor.Item{
			Type: deepcolor.ItemTypeMapList,
			Rules: []deepcolor.ItemRule{
				{
					Key:      "title",
					Selector: ".result-item-title.result-game-item-title > a > span",
					Target:   deepcolor.TextTarget(),
				},
				{
					Selector: ".result-item.result-game-item > .result-game-item-pic > a",
					Key:      "url",
					Target:   deepcolor.AttributeTarget("href"),
				},
				{
					Selector: ".result-item.result-game-item > .result-game-item-pic > a > img",
					Key:      "cover",
					Target:   deepcolor.AttributeTarget("src"),
				},
				{
					Key:      "abstraction",
					Selector: ".result-game-item-detail > .result-game-item-desc",
					Target:   deepcolor.TextTarget(),
				},
				{
					Key:      "author",
					Selector: ".result-game-item-detail > .result-game-item-info > p:nth-of-type(1) > span::nth-of-type(2)",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
	}
	BiqugeBRule = NovelProviderRule{
		Title: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#info > h1",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Author: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#info > p:nth-of-type(1)",
					Target:   deepcolor.TextTarget(),
					Filters:  []string{"作(.*)者:"},
				},
			},
		},
		Cover: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#fmimg > img",
					Target:   deepcolor.AttributeTarget("src"),
				},
			},
		},
		Abstraction: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#intro",
					Target:   deepcolor.TextTarget(),
					Filters:  []string{"本站提示：各位书友要是觉得(.*)还不错的话请不要忘记向您QQ群和微博里的朋友推荐哦！"},
				},
			},
		},
		Chapters: deepcolor.Item{
			Type: deepcolor.ItemTypeMapList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#list > dl > dd > a",
					Key:      "name",
					Target:   deepcolor.TextTarget(),
				},
				{
					Selector: "#list > dl > dd > a",
					Key:      "url",
					Target:   deepcolor.AttributeTarget("href"),
				},
			},
		},
		Content: deepcolor.Item{
			Type: deepcolor.ItemTypeMap,
			Rules: []deepcolor.ItemRule{
				{
					Key:      "content",
					Selector: "#content",
					Target:   deepcolor.TextTarget(),
					Filters:  []string{"\u00a0"},
				},
				{
					Key:      "name",
					Selector: ".bookname > h1",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Search: deepcolor.Item{
			Type: deepcolor.ItemTypeMapList,
			Rules: []deepcolor.ItemRule{
				{
					Key:      "title",
					Selector: "tbody > tr[align!=center] > td:nth-of-type(1) > a",
					Target:   deepcolor.TextTarget(),
				},
				{
					Key:      "url",
					Selector: "tbody > tr[align!=center] > td:nth-of-type(1) > a",
					Target:   deepcolor.AttributeTarget("href"),
				},
				{
					Key:      "cover",
					Selector: "",
					Target:   deepcolor.TextTarget(),
				},
				{
					Key:      "abstraction",
					Selector: "",
					Target:   deepcolor.TextTarget(),
				},
				{
					Key:      "author",
					Selector: "tbody > tr[align!=center] > td:nth-of-type(3)",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
	}
	BiqugeCRule = NovelProviderRule{
		Title: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".divbox.cf > div:nth-of-type(2) > div:nth-of-type(1) > span:nth-of-type(1)",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Author: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".divbox.cf > div:nth-of-type(2) > div:nth-of-type(1) > span:nth-of-type(2) > a",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Cover: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".divbox.cf > div:nth-of-type(1) > a",
					Target:   deepcolor.AttributeTarget("href"),
				},
			},
		},
		Abstraction: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".tabcontent > .tabvalue > div",
					Target:   deepcolor.TextTarget(),
					Filters:  []string{"(.*)的其它作品", "\u2002", "\u2003"},
				},
			},
		},
		Chapters: deepcolor.Item{
			Type: deepcolor.ItemTypeMapList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".main > div:nth-of-type(3) > .divbg > .infoindex > dd > a",
					Key:      "name",
					Target:   deepcolor.TextTarget(),
				},
				{
					Selector: ".main > div:nth-of-type(3) > .divbg > .infoindex > dd > a",
					Key:      "url",
					Target:   deepcolor.AttributeTarget("href"),
				},
			},
		},
		Content: deepcolor.Item{
			Type: deepcolor.ItemTypeMap,
			Rules: []deepcolor.ItemRule{
				{
					Key:      "content",
					Selector: "#acontent",
					Target:   deepcolor.TextTarget(),
					Filters:  []string{"笔趣阁", "\u2002", "(.*)，最快更新(.*)最新章节！", "\u2003"},
				},
				{
					Key:      "name",
					Selector: ".atitle",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Search: deepcolor.Item{
			Type: deepcolor.ItemTypeMapList,
			Rules: []deepcolor.ItemRule{
				{
					Key:      "title",
					Selector: "#jieqi_page_contents > .c_row > div:nth-of-type(2) >div:nth-of-type(1) > span > a > span",
					Target:   deepcolor.TextTarget(),
				},
				{
					Key:      "url",
					Selector: "#jieqi_page_contents > .c_row > div:nth-of-type(2) >div:nth-of-type(1) > span > a",
					Target:   deepcolor.AttributeTarget("href"),
				},
				{
					Selector: "#jieqi_page_contents > .c_row > .fl > a > img",
					Key:      "cover",
					Target:   deepcolor.AttributeTarget("src"),
				},
				{
					Key:      "abstraction",
					Selector: ".c_description",
					Target:   deepcolor.TextTarget(),
				},
				{
					Key:      "author",
					Selector: "#jieqi_page_contents > .c_row > div:nth-of-type(2) >div:nth-of-type(2) > span:nth-of-type(2)",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
	}

	LigntNovelRule = NovelProviderRule{
		Title: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".book-name",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Author: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".au-name > a",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Cover: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".book-img > img",
					Target:   deepcolor.AttributeTarget("src"),
				},
			},
		},
		Abstraction: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".book-dec > p",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		ChapaterUrl: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".btn.read-btn",
					Target:   deepcolor.AttributeTarget("href"),
				},
			},
		},
		Chapters: deepcolor.Item{
			Type: deepcolor.ItemTypeMapList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".chapter-list > li > a",
					Key:      "name",
					Target:   deepcolor.TextTarget(),
				},
				{
					Selector: ".chapter-list > li > a",
					Key:      "url",
					Target:   deepcolor.AttributeTarget("href"),
				},
			},
		},
		Content: deepcolor.Item{
			Type: deepcolor.ItemTypeMap,
			Rules: []deepcolor.ItemRule{
				{
					Key:      "content",
					Selector: "#TextContent",
					Target:   deepcolor.TextTarget(),
					Filters:  []string{"style_bm\\(\\);", "style_tp\\(\\);"},
				},
				{
					Key:      "name",
					Selector: "#mlfy_main_text > h1",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Search: deepcolor.Item{
			Type: deepcolor.ItemTypeMapList,
			Rules: []deepcolor.ItemRule{
				{
					Key:      "title",
					Selector: ".search-result-list > .se-result-infos > h2 > a",
					Target:   deepcolor.TextTarget(),
				},
				{
					Key:      "url",
					Selector: ".search-result-list > .se-result-infos > h2 > a",
					Target:   deepcolor.AttributeTarget("href"),
				},
				{
					Key:      "cover",
					Selector: ".search-result-list > .imgbox > a > img",
					Target:   deepcolor.AttributeTarget("src"),
				},
				{
					Key:      "abstraction",
					Selector: ".search-result-list > .se-result-infos > p",
					Target:   deepcolor.TextTarget(),
				},
				{
					Key:      "author",
					Selector: ".search-result-list > .se-result-infos > .bookinfo > a:nth-of-type(1)",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
	}
)

var (
	BiqugeProvider = NovelProvider{
		Identifier: "biqugeA",
		Name:       "笔趣阁 A",
		Alias:      "biquege.com.cn",
		HomeUrl:    "https://www.biquge.com.cn/",
		Charset:    "utf-8",
		Status:     true,
		InfoUrl:    "^http(s)?://www.biquge.com.cn/book/[0-9]+(/)?$",
		ContentUrl: "^https(:)?//www.biquge.com.cn/book/[0-9]+/[0-9]+\\.html$",
		SearchApi:  "https://www.biquge.com.cn/search.php?q=%s",
		Rule:       BiqugeARule,
	}
	BiqugeBProvider = NovelProvider{
		Identifier: "biqugeB",
		Name:       "笔趣阁 B",
		Alias:      "biquwx.la",
		HomeUrl:    "https://www.biquwx.la/",
		Charset:    "utf-8",
		Status:     true,
		InfoUrl:    "^http(s)?://www.biquwx.la/[0-9]+_[0-9]+(/)?$",
		ContentUrl: "^http(s)?://www.biquwx.la/[0-9]+_[0-9]+/[0-9]+\\.html$",
		SearchApi:  "https://www.biquwx.la/modules/article/search.php?searchkey=%s",
		Rule:       BiqugeBRule,
	}
	BiqugeCProvider = NovelProvider{
		Identifier: "biqugeC",
		Name:       "笔趣阁 B",
		Alias:      "biquwx.la",
		HomeUrl:    "http://www.liquge.com/",
		Charset:    "utf-8",
		Status:     true,
		InfoUrl:    "^http(s)?://www.liquge.com/book/[0-9]+(/)?$",
		ContentUrl: "^http(s)?://www.liquge.com/book/[0-9]+/[0-9]+\\.html$",
		SearchApi:  "http://www.liquge.com/modules/article/search.php?searchkey=%s",
		Rule:       BiqugeCRule,
	}
	LigntNovelProvider = NovelProvider{
		Identifier: "linovelib",
		Name:       "哔哩轻小说",
		Alias:      "linovelib.com",
		HomeUrl:    "https://www.linovelib.com/",
		Charset:    "utf-8",
		Header: map[string]string{
			"cookie": "PHPSESSID=47c8s8ti7n0j3nk10l1mjl0g6n;",
		},
		Status:     true,
		InfoUrl:    "^http(s)?://www.linovelib.com/novel/[0-9]+\\.html$",
		ContentUrl: "^http(s)?://www.linovelib.com/novel/[0-9]+/[0-9]+\\.html$",
		SearchApi:  "https://www.linovelib.com/s/?searchkey=%s",
		Rule:       LigntNovelRule,
	}
)
