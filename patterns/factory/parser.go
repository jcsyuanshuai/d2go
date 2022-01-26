package factory

type RuleParser interface {
	Parse(data []byte)
}

type jsonRuleParser struct {
}

func (j jsonRuleParser) Parse(data []byte) {

}

type yamlRuleParser struct {
}

func (y yamlRuleParser) Parse(data []byte) {
}

func NewRuleParser(name string) RuleParser {
	switch name {
	case "json":
		return jsonRuleParser{}
	case "yaml":
		return yamlRuleParser{}
	}
	return nil
}
