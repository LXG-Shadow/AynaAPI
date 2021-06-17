package deepcolor

type RuleType string

const (
	RuleTypeSingleItem  RuleType = "single"
	RuleTypeListItem    RuleType = "list"
	RuleTypeMapItem     RuleType = "map"
	RuleTypeMapListItem RuleType = "maplist"
)

type SelectorTargetType string

const (
	SelectorTargetTypeInnerText SelectorTargetType = "innertext"
	SelectorTargetTypeAttribute SelectorTargetType = "attribute"
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

type RuleCollection struct {
	Type  RuleType
	Rules []ItemRule
}

type ItemRule struct {
	Key      string
	Selector string
	Target   SelectorTarget
	Filters  []string
}
