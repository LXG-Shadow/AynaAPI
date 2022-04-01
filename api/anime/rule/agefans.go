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

	InfoVideos     deepcolor.Item
	InfoVideoNames deepcolor.Item
}

func InitializeAgefansRules() AgefansRules {
	return AgefansRules{
		SearchURL: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.AttributeSelector(
						"a.cell_imform_name",
						"href"),
				},
			},
		},
		SearchCover: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.AttributeSelector("a.cell_poster > img",
						"src",
					),
				},
			},
		},
		SearchTitle: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector(
						"a.cell_imform_name"),
				},
			},
		},
		SearchYear: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector(
						"div.blockcontent1 > div.cell > div.cell_imform > div.cell_imform_kvs > div:nth-child(4) > span.cell_imform_value",
					),
				},
			},
		},
		SearchTag: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector(
						"div.blockcontent1 > div.cell > div.cell_imform > div.cell_imform_kvs > div:nth-child(7) > span.cell_imform_value"),
				},
			},
		},
		SearchDesc: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector("div.cell_imform_desc"),
				},
			},
		},

		InfoTitle: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector(".detail_imform_name"),
				},
			},
		},
		InfoCover: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.AttributeSelector(
						"img.poster",
						"src"),
				},
			},
		},
		InfoYear: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector("ul.blockcontent > li.detail_imform_kv:nth-child(7) > span.detail_imform_value"),
				},
			},
		},
		InfoTag: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector(
						"ul.blockcontent > li.detail_imform_kv:nth-child(10) > span.detail_imform_value"),
				},
			},
		},
		InfoDesc: deepcolor.Item{
			Type: deepcolor.ItemTypeSingle,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector("div.detail_imform_desc_pre > p"),
				},
			},
		},
		InfoVideos: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.AttributeSelector(
						"div.movurl > ul > li > a",
						"href",
					),
					Substitution: map[string]string{
						"/play/":     "",
						"\\?playid=": "-",
						"_":          "-",
					},
				},
			},
		},
		InfoVideoNames: deepcolor.Item{
			Type: deepcolor.ItemTypeList,
			Rules: []deepcolor.ItemRule{
				{
					Selector: deepcolor.TextSelector("div.movurl > ul > li > a"),
				},
			},
		},
	}
}
