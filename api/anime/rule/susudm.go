package rule

import "github.com/aynakeya/deepcolor"

type SusuDmRules struct {
	Title     deepcolor.Item
	Cover     deepcolor.Item
	Tags      deepcolor.Item
	Desc      deepcolor.Item
	AreaYear  deepcolor.Item
	AreaYear2 deepcolor.Item
}

func InitializeSusuDmRules() SusuDmRules {
	return SusuDmRules{
		Title: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "dt.name",
					Target:   deepcolor.TextTarget(),
					Substitution: map[string]string{
						"<span[^>]*>.*</span>": "",
					},
				},
			},
		},
		Cover: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "body > div.wrap > div.content.mb.clearfix > div.pic > img",
					Target:   deepcolor.AttributeTarget("src"),
				},
			},
		},
		Tags: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "body > div.wrap > div.content.mb.clearfix > div.info > dl > dd:nth-child(4) > a",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		Desc: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "div.des2",
					Target:   deepcolor.TextTarget(),
					Substitution: map[string]string{
						"剧情：": "",
					},
				},
			},
		},
		AreaYear: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "body > div.wrap > div.content.mb.clearfix > div.info > dl > dd:nth-child(3)",
					Target:   deepcolor.TextTarget(),
					Substitution: map[string]string{
						"<b>地区：</b>":       "",
						"(\\s)*<b>年代：</b>": "---",
					},
				},
			},
		},
		AreaYear2: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "body > div.wrap > div.content.mb.clearfix > div.info > dl > dd:nth-child(2)",
					Target:   deepcolor.TextTarget(),
					Substitution: map[string]string{
						"<b>地区：</b>":       "",
						"(\\s)*<b>年代：</b>": "---",
					},
				},
			},
		},
	}
}
