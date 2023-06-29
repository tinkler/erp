package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/tinkler/erp/test"
)

func TestGetUser(t *testing.T) {
	test.LoadEnv()
	admin := &User{Username: "admin"}
	if err := admin.Get(test.Context()); err != nil {
		t.Fatal(err)
	}
	if _, err := uuid.Parse(admin.ID); err != nil {
		t.Fatal(err)
	}
	if admin.PhoneNumber != "18176386025" {
		t.Fail()
	}
	if len(admin.Emails) != 0 {
		t.Fail()
	}
	admin = &User{ID: admin.ID}
	if err := admin.GetDetail(test.Context()); err != nil {
		t.Fatal(err)
	}
	if len(admin.Emails) == 0 {
		t.Fatal()
	}
	if admin.Emails[0].Address != "tinkler@163.com" {
		t.Fail()
	}
}
