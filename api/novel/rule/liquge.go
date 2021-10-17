package rule

import "github.com/aynakeya/deepcolor"

func InitializeLiqugeRules() BiqugeRule {
	return BiqugeRule{Title: deepcolor.Item{
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
		Description: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".tabcontent > .tabvalue > div",
					Target:   deepcolor.TextTarget(),
					Substitution: map[string]string{
						"(.*)的其它作品": "",
						"\u2002":    "",
						"\u2003":    "",
					},
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
					Substitution: map[string]string{
						"笔趣阁":                "",
						"\u2002":             "",
						"(.*)，最快更新(.*)最新章节！": "",
						"\u2003":             "",
					},
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
					Key:      "description",
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
}
