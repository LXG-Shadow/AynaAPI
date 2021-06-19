package deepcolor

type ItemType string

const (
	ItemTypeSingle  ItemType = "single"
	ItemTypeList    ItemType = "list"
	ItemTypeMap     ItemType = "map"
	ItemTypeMapList ItemType = "maplist"
)

type SelectorTargetType string

const (
	SelectorTargetTypeInnerText SelectorTargetType = "innertext"
	SelectorTargetTypeAttribute SelectorTargetType = "attribute"
	SelectorTargetTypeRegExp    SelectorTargetType = "regexp"

	SelectorTargetTypeJson SelectorTargetType = "json"
)

type SelectorTarget struct {
	Type  SelectorTargetType
	Value string
}

func TextTarget() SelectorTarget {
	return SelectorTarget{SelectorTargetTypeInnerText, ""}
}

func AttributeTarget(attribute string) SelectorTarget {
	return SelectorTarget{SelectorTargetTypeAttribute, attribute}
}

func RegularExpressionTarget() SelectorTarget {
	return SelectorTarget{SelectorTargetTypeRegExp, ""}
}

func JsonTarget() SelectorTarget {
	return SelectorTarget{SelectorTargetTypeJson, ""}
}

type Item struct {
	Type  ItemType   `json:"type"`
	Rules []ItemRule `json:"rules"`
}

type ItemRule struct {
	Key       string            `json:"key"`
	Selector  string            `json:"selector"`
	Target    SelectorTarget    `json:"target"`
	Filters   []string          `json:"filters"`
	Replacers map[string]string `json:"replacers"`
}
