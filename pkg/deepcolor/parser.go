package deepcolor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"regexp"
)

func applyFilters(content string, rule ItemRule) string {
	if rule.Filters == nil {
		return content
	}
	for _, f := range rule.Filters {
		re := regexp.MustCompile(f)
		content = re.ReplaceAllString(content, "")
	}
	return content
}

func getValue(selection *goquery.Selection, rule ItemRule) string {
	if rule.Selector == "" {
		return ""
	}
	switch rule.Target.Type {
	case SelectorTargetTypeInnerText:
		return applyFilters(selection.Text(), rule)
	case SelectorTargetTypeAttribute:
		attr, _ := selection.Attr(rule.Target.Value)
		return applyFilters(attr, rule)
	default:
		return applyFilters(selection.Text(), rule)
	}
}

func ParseSingle(doc *goquery.Document, collection Item) (result string) {
	result = ""
	if collection.Type != ItemTypeSingle {
		return
	}
	if len(collection.Rules) != 1 {
		return
	}
	result = getValue(doc.Find(collection.Rules[0].Selector), collection.Rules[0])
	return
}

func ParseList(doc *goquery.Document, collection Item) (result []string) {
	result = make([]string, 0)
	if collection.Type != ItemTypeList {
		return
	}
	if len(collection.Rules) != 1 {
		return
	}
	rule := collection.Rules[0]
	doc.Find(rule.Selector).Each(func(i int, selection *goquery.Selection) {
		result = append(result, getValue(selection, rule))
	})
	return
}

func ParseMap(doc *goquery.Document, collection Item) (result map[string]string) {
	result = map[string]string{}
	if collection.Type != ItemTypeMap {
		return
	}
	for _, rule := range collection.Rules {
		result[rule.Key] = getValue(doc.Find(rule.Selector), rule)
	}
	return
}

func newBaseMap(collection Item) map[string]string {
	baseMap := map[string]string{}
	for _, rule := range collection.Rules {
		baseMap[rule.Key] = ""
	}
	return baseMap
}

func ParseMapList(doc *goquery.Document, collection Item) (result []map[string]string) {
	result = make([]map[string]string, 0)
	if collection.Type != ItemTypeMapList {
		return
	}
	for _, rule := range collection.Rules {
		doc.Find(rule.Selector).Each(func(index int, selection *goquery.Selection) {
			if len(result) <= index {
				result = append(result, newBaseMap(collection))
			}
			result[index][rule.Key] = getValue(selection, rule)
		})
	}
	return
}

func ParseJsonSingle(doc *gjson.Result, item Item) (result string) {
	result = ""
	if item.Type != ItemTypeSingle {
		return
	}
	if len(item.Rules) != 1 {
		return
	}
	result = doc.Get(item.Rules[0].Selector).String()
	return
}

func ParseJsonList(doc *gjson.Result, collection Item) (result []string) {
	result = make([]string, 0)
	if collection.Type != ItemTypeList {
		return
	}
	if len(collection.Rules) != 1 {
		return
	}
	rule := collection.Rules[0]
	doc.Get(rule.Selector).ForEach(func(key, value gjson.Result) bool {
		result = append(result, value.String())
		return true
	})
	return
}

func ParseJsonMap(doc *gjson.Result, collection Item) (result map[string]string) {
	result = map[string]string{}
	if collection.Type != ItemTypeMap {
		return
	}
	for _, rule := range collection.Rules {
		result[rule.Key] = doc.Get(rule.Selector).String()
	}
	return
}

func ParseJsonMapList(doc *gjson.Result, collection Item) (result []map[string]string) {
	result = make([]map[string]string, 0)
	if collection.Type != ItemTypeMapList {
		return
	}
	for _, rule := range collection.Rules {
		index := 0
		doc.Get(rule.Selector).ForEach(func(key, value gjson.Result) bool {
			if len(result) <= index {
				result = append(result, newBaseMap(collection))
			}
			result[index][rule.Selector] = value.String()
			return true
		})
	}
	return
}
