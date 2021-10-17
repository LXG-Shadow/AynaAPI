package rule

import "github.com/aynakeya/deepcolor"

func InitializeBiquwxRules() BiqugeRule {
	return BiqugeRule{
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
					Selector:     "#info > p:nth-of-type(1)",
					Target:       deepcolor.TextTarget(),
					Substitution: map[string]string{"作(.*)者:": ""},
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
		Description: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "#intro",
					Target:   deepcolor.TextTarget(),
					Substitution: map[string]string{
						"本站提示：各位书友要是觉得(.*)还不错的话请不要忘记向您QQ群和微博里的朋友推荐哦！": "",
					},
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
					Key:          "content",
					Selector:     "#content",
					Target:       deepcolor.TextTarget(),
					Substitution: map[string]string{"\u00a0": ""},
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
					Key:      "description",
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
}
