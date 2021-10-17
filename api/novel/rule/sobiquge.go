package rule

import "github.com/aynakeya/deepcolor"

type BiqugeRule struct {
	Title       deepcolor.Item `json:"title"`
	Author      deepcolor.Item `json:"author"`
	Cover       deepcolor.Item `json:"cover"`
	Description deepcolor.Item `json:"description"`
	Chapters    deepcolor.Item `json:"chapters"`
	Content     deepcolor.Item `json:"content"`
	Search      deepcolor.Item `json:"search"`
}

func InitializeSobiqugeRules() BiqugeRule {
	return BiqugeRule{
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
					Selector:     "#info > p:nth-of-type(1)",
					Target:       deepcolor.TextTarget(),
					Substitution: map[string]string{"作(.*)者：": ""},
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
					Key:      "description",
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
}
