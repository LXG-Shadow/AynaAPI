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

func applyReplacers(content string, rule ItemRule) string {
	if rule.Replacers == nil {
		return content
	}
	for key, val := range rule.Replacers {
		content = regexp.MustCompile(key).ReplaceAllString(content, val)
	}
	return content
}

func applySelectionReplacers(selection *goquery.Selection, rule ItemRule) *goquery.Selection {
	if rule.Replacers == nil {
		return selection
	}
	htmltext, _ := selection.Html()
	for key, val := range rule.Replacers {
		htmltext = regexp.MustCompile(key).ReplaceAllString(htmltext, val)
	}
	return selection.SetHtml(htmltext)
}

func getValue(selection *goquery.Selection, rule ItemRule) string {
	if rule.Selector == "" {
		return ""
	}
	selection = applySelectionReplacers(selection, rule)
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
	if len(collection.Rules) < 1 {
		return
	}
	for _, rule := range collection.Rules {
		switch rule.Target.Type {
		case SelectorTargetTypeRegExp:
			htext, _ := doc.Html()
			result += applyFilters(applyReplacers(regexp.MustCompile(rule.Selector).FindString(htext), rule), rule)
		default:
			result += getValue(doc.Find(rule.Selector), rule)
		}
	}
	return
}

func ParseList(doc *goquery.Document, collection Item) (result []string) {
	result = make([]string, 0)
	if collection.Type != ItemTypeList {
		return
	}
	if len(collection.Rules) < 1 {
		return
	}
	for _, rule := range collection.Rules {
		doc.Find(rule.Selector).Each(func(i int, selection *goquery.Selection) {
			if len(result) <= i {
				result = append(result, getValue(selection, rule))
			} else {
				result[i] += getValue(selection, rule)
			}
		})
	}
	return
}

func ParseMap(doc *goquery.Document, collection Item) (result map[string]string) {
	result = map[string]string{}
	if collection.Type != ItemTypeMap {
		return
	}
	for _, rule := range collection.Rules {
		switch rule.Target.Type {
		case SelectorTargetTypeRegExp:
			htext, _ := doc.Html()
			result[rule.Key] += applyFilters(applyReplacers(regexp.MustCompile(rule.Selector).FindString(htext), rule), rule)
		default:
			result[rule.Key] += getValue(doc.Find(rule.Selector), rule)
		}
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
		switch rule.Target.Type {
		case SelectorTargetTypeRegExp:
			htext, _ := doc.Html()
			for index, val := range regexp.MustCompile(rule.Selector).FindAllString(htext, 0) {
				if len(result) <= index {
					result = append(result, newBaseMap(collection))
				}
				result[index][rule.Key] += applyFilters(applyReplacers(val, rule), rule)
			}
		default:
			doc.Find(rule.Selector).Each(func(index int, selection *goquery.Selection) {
				if len(result) <= index {
					result = append(result, newBaseMap(collection))
				}
				result[index][rule.Key] += getValue(selection, rule)
			})
		}
	}
	return
}

// todo json same key append

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
