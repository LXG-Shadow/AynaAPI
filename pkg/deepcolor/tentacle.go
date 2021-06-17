package deepcolor

import (
	"github.com/PuerkitoBio/goquery"
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

func Parse(doc *goquery.Document, collection RuleCollection) interface{} {
	switch collection.Type {
	case RuleTypeSingleItem:
		return ParseSingle(doc, collection)
	case RuleTypeListItem:
		return ParseList(doc, collection)
	case RuleTypeMapItem:
		return ParseMap(doc, collection)
	case RuleTypeMapListItem:
		return ParseMapList(doc, collection)
	default:
		return ""
	}
}

func ParseSingle(doc *goquery.Document, collection RuleCollection) (result string) {
	result = ""
	if collection.Type != RuleTypeSingleItem {
		return
	}
	if len(collection.Rules) != 1 {
		return
	}
	result = getValue(doc.Find(collection.Rules[0].Selector), collection.Rules[0])
	return
}

func ParseList(doc *goquery.Document, collection RuleCollection) (result []string) {
	result = make([]string, 0)
	if collection.Type != RuleTypeListItem {
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

func ParseMap(doc *goquery.Document, collection RuleCollection) (result map[string]string) {
	result = map[string]string{}
	if collection.Type != RuleTypeMapItem {
		return
	}
	for _, rule := range collection.Rules {
		result[rule.Key] = getValue(doc.Find(rule.Selector), rule)
	}
	return
}

func newBaseMap(collection RuleCollection) map[string]string {
	baseMap := map[string]string{}
	for _, rule := range collection.Rules {
		baseMap[rule.Key] = ""
	}
	return baseMap
}

func ParseMapList(doc *goquery.Document, collection RuleCollection) (result []map[string]string) {
	result = make([]map[string]string, 0)
	if collection.Type != RuleTypeMapListItem {
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
