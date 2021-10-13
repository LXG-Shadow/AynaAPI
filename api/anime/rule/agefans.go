package rule

import "github.com/aynakeya/deepcolor"

type AgefansRules struct {
	SearchURL   deepcolor.Item
	SearchTitle deepcolor.Item
	SearchCover deepcolor.Item
	SearchYear  deepcolor.Item
	SearchTag   deepcolor.Item
	SearchDesc  deepcolor.Item

	InfoTitle deepcolor.Item
	InfoCover deepcolor.Item
	InfoYear  deepcolor.Item
	InfoTag   deepcolor.Item
	InfoDesc  deepcolor.Item
}

func InitializeAgefansRules() AgefansRules {
	return AgefansRules{
		SearchURL: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "a.cell_imform_name",
					Target:   deepcolor.AttributeTarget("href"),
				},
			},
		},
		SearchCover: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "a.cell_poster > img",
					Target:   deepcolor.AttributeTarget("src"),
				},
			},
		},
		SearchTitle: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "a.cell_imform_name",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		SearchYear: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "div.blockcontent1 > div.cell > div.cell_imform > div.cell_imform_kvs > div:nth-child(4) > span.cell_imform_value",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		SearchTag: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "div.blockcontent1 > div.cell > div.cell_imform > div.cell_imform_kvs > div:nth-child(7) > span.cell_imform_value",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		SearchDesc: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "div.cell_imform_desc",
					Target:   deepcolor.TextTarget(),
				},
			},
		},

		InfoTitle: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: ".detail_imform_name",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		InfoCover: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "img.poster",
					Target:   deepcolor.AttributeTarget("src"),
				},
			},
		},
		InfoYear: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "ul.blockcontent > li.detail_imform_kv:nth-child(7) > span.detail_imform_value",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		InfoTag: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "ul.blockcontent > li.detail_imform_kv:nth-child(10) > span.detail_imform_value",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
		InfoDesc: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: "div.detail_imform_desc_pre > p",
					Target:   deepcolor.TextTarget(),
				},
			},
		},
	}
}
