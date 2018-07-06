package models

import "testing"

func Test_UserPermission(t *testing.T) {

	var u User
	u.Username = "user-perm-fido"

	var p Permission
	p.Name = "Moms paghetti"
	p.Description = "xd"

	if errs, err := db.ValidateAndCreate(&u); err != nil || errs.HasAny() {
		t.Error(errs)
		t.Error(err)
	}

	if errs, err := db.ValidateAndCreate(&p); err != nil || errs.HasAny() {
		t.Error(errs)
		t.Error(err)
	}

	var uperm UserPermission
	uperm.UserID = u.ID
	uperm.PermissionID = p.ID

	if errs, err := db.ValidateAndCreate(&uperm); err != nil || errs.HasAny() {
		t.Error(errs)
		t.Error(err)
	}

	if err := db.Load(&p); err != nil {
		t.Error(err)
	}

	if err := db.Load(&u); err != nil {
		t.Error(err)
	}

	if p.Accessors[0].ID != u.ID {
		t.Error("Expected user id")
	}

	if u.Perms[0].ID != p.ID {
		t.Error("Expected perm id")
	}

}
