package commodity

import (
	"errors"
	"net/http"
	"testing"

	"github.com/tinkler/erp/mist/pssm/test"
	"github.com/tinkler/mqttadmin/pkg/status"
)

func assertErr(t *testing.T, err error, target error) {
	if !errors.Is(err, target) {
		t.Fatal(target)
	}
}

func testDoubleCreate(t *testing.T, m *Commodity) {
	m.ID = ""
	err := m.Create(test.Context())
	if err != nil {
		t.Fatal(err)
	}
	// TODO check if commodity is created twice
}

func TestCreate(t *testing.T) {
	test.LoadEnv()
	ctx := test.Context()
	m := new(Commodity)
	m.ID = "1"
	err := m.Create(ctx)
	assertErr(t, status.NewCn(http.StatusBadRequest, "ID is not empty", "新增商品ID不为空"), err)
	m.ID = ""
	err = m.Create(ctx)
	assertErr(t, status.NewCn(http.StatusBadRequest, "Name is empty", "新增商品名称为空"), err)
	// name
	m.Name = "家家面"
	err = m.Create(ctx)
	assertErr(t, status.NewCn(http.StatusBadRequest, "Barcode is empty", "新增商品二维码为空"), err)
	m.Barcode = "695231203999"
	err = m.Create(ctx)
	assertErr(t, status.NewCn(http.StatusBadRequest, "Unit is empty", "新增商品单位为空"), err)
	m.Unit = "箱"
	err = m.Create(ctx)
	assertErr(t, status.NewCn(http.StatusBadRequest, "Category is empty", "新增商品类别为空"), err)
	m.Category = "正麦"
	err = m.Create(ctx)
	assertErr(t, nil, err)
	if m.ID == "" {
		t.Fatal("Commodity failed to create")
	}
	testDoubleCreate(t, m)
}
