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

func Test_UserPermission_UniqueEntry(t *testing.T) {

	var u, u2 = &User{}, &User{}
	var p = &Permission{}

	u.Username = "pepep-testus"
	p.Name = "permiso-pep"
	p.Description = "permiso de pepep"

	u2.Username = "pepo-testus"

	if _, err := db.ValidateAndCreate(u); err != nil {
		t.Error(err)
	}

	if _, err := db.ValidateAndCreate(u2); err != nil {
		t.Error(err)
	}

	if _, err := db.ValidateAndCreate(p); err != nil {
		t.Error(err)
	}

	var up, dup, up2 = &UserPermission{}, &UserPermission{}, &UserPermission{}

	up.UserID = u.ID
	up.PermissionID = p.ID

	dup.UserID = u.ID
	dup.PermissionID = p.ID

	up2.UserID = u2.ID
	up2.PermissionID = p.ID

	if errs, err := db.ValidateAndCreate(up); err != nil || errs.HasAny() {
		t.Error(errs)
		t.Error(err)
	}

	if errs, err := db.ValidateAndCreate(dup); err == nil && !errs.HasAny() {
		t.Error("missed expected duplication error")
	}

	if errs, err := db.ValidateAndCreate(up2); err != nil || errs.HasAny() {
		t.Error(errs)
		t.Error(err)
	}
}

//TODO
func Test_UserPermission_RemovePerms(t *testing.T) {
	t.Fatal("unimplemented")
}
