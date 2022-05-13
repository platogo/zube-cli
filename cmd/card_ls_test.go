package cmd

import (
	"reflect"
	"testing"

	"github.com/platogo/zube-cli/utils"
	"github.com/platogo/zube-cli/zube"
)

func TestNewQueryFromFlags(t *testing.T) {
	example := cardLsCmd.Flags()

	example.Set("category", "Inbox")
	example.Set("priority", "3")
	example.Set("status", "open")
	example.Set("number", "0")

	res := utils.NewQueryFromFlags(example)

	want := zube.Query{Filter: zube.Filter{
		Where:  map[string]any{"category_name": "Inbox", "priority": 3, "status": "open"},
		Select: []string{"number", "title", "status", "category_name"}},
	}

	if !reflect.DeepEqual(res, want) {
		t.Errorf("want does not match result, \nexpected: %+v \ngot: %+v", want, res)
	}
}
