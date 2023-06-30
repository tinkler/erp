package inventory

import (
	"testing"

	"github.com/tinkler/erp/mist/pssm/test"
)

func TestCreateStoreIn(t *testing.T) {
	test.LoadEnv()
	ctx := test.Context()

	m := StoreIn{
		StoreInItems: []*StoreInItem{
			{
				CommodityID: "381b8343-c7b2-48ff-b4dd-00a7b56a6078",
				Quantity:    10,
			},
		},
	}

	err := m.Create(ctx)
	if err != nil {
		t.Fatal(err)
	}

}
