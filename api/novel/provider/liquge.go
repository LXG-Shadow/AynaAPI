package provider

import "AynaAPI/api/novel/rule"

var LiqugeAPI *Biquge

func __newLiquge() *Biquge {
	return &Biquge{
		Name:      "liquge",
		BaseUrl:   "http://www.liquge.com/",
		Charset:   "utf-8",
		SearchAPI: "http://www.liquge.com/modules/article/search.php?searchkey=%s",
		Rules:     rule.InitializeLiqugeRules(),
	}
}

func init() {
	LiqugeAPI = __newLiquge()
}
