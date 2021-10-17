package rule

import "github.com/aynakeya/deepcolor"

type LinovelibRule struct {
	Title        deepcolor.Item `json:"title"`
	Author       deepcolor.Item `json:"author"`
	Cover        deepcolor.Item `json:"cover"`
	Description  deepcolor.Item `json:"description"`
	Chapters     deepcolor.Item `json:"chapters"`
	ChapterEntry deepcolor.Item `json:"chapter_entry"`
	Content      deepcolor.Item `json:"content"`
	ContentNext  deepcolor.Item `json:"content_next"`
	Search       deepcolor.Item `json:"search"`
}

func InitializeLinovelibRules() LinovelibRule {
	return LinovelibRule{
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
		Description: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".book-dec > p",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		ChapterEntry: deepcolor.Item{
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
					Substitution: map[string]string{
						"style_bm\\(\\);":       "",
						"style_tp\\(\\);":       "",
						"\\n\\n":                "",
						"</p>":                  "\n</p>",
						"<span[^>]*>.*?</span>": "",
					},
				},
				//{
				//	Key:      "content",
				//	Selector: "let dom_nr = '(.*)';document",
				//	Target:   deepcolor.RegularExpressionTarget(),
				//	Filters: []string{"let dom_nr = '", "';document", "（本章未完）",
				//		"<!-- <p", "<p[^>]*>", "</p>", "<!-- ", " -->", "<img[^>]*>"},
				//	Replacers: map[string]string{
				//		"</p>": "\n</p>",
				//	},
				//},
				//{
				//	Key:      "name",
				//	Selector: "#mlfy_main_text > h1",
				//	Target:   deepcolor.TextTarget(),
				//},
			},
		},
		ContentNext: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".mlfy_page > a:nth-of-type(5)",
					Target:   deepcolor.AttributeTarget("href"),
					Substitution: map[string]string{
						"/novel/[0-9]+/[0-9]+\\.html": "",
					},
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
					Key:      "description",
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
}
