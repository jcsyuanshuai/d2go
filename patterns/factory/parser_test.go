package factory

import (
	"reflect"
	"testing"
)

type Item struct {
	name   string
	args   string
	expect RuleParser
}

func TestNewRuleParser(t *testing.T) {
	cases := [2]Item{
		{
			name:   "json",
			args:   "json",
			expect: jsonRuleParser{},
		}, {
			name:   "yaml",
			args:   "yaml",
			expect: yamlRuleParser{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if p := NewRuleParser(c.args); !reflect.DeepEqual(p, c.expect) {
				t.Errorf("NewRuleParser() = %v, expect %v", p, c.expect)
			}
		})
	}
}
