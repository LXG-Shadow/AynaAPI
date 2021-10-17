package provider

import "AynaAPI/api/novel/rule"

var BiquwxAPI *Biquge

func __newBiquwx() *Biquge {
	return &Biquge{
		Name:      "biquwx",
		BaseUrl:   "https://www.biquwx.la/",
		Charset:   "utf-8",
		SearchAPI: "https://www.biquwx.la/modules/article/search.php?searchkey=%s",
		Rules:     rule.InitializeBiquwxRules(),
	}
}

func init() {
	BiquwxAPI = __newBiquwx()
}
