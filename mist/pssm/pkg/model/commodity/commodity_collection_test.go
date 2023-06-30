package commodity

import (
	"testing"

	"github.com/tinkler/erp/mist/pssm/test"
)

func TestSearchByName(t *testing.T) {
	test.LoadEnv()
	ctx := test.Context()
	c := new(CommodityCollection)
	c.SearchByName(ctx, "家家面")
	if len(c.List) == 0 {
		t.Fatal()
	}
}

func TestSearchByCode(t *testing.T) {
	test.LoadEnv()
	ctx := test.Context()
	c := new(CommodityCollection)
	c.SearchByCode(ctx, "%3999")
	if len(c.List) == 0 {
		t.Fail()
	}
}
