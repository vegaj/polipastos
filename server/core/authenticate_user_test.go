package core

import (
	"testing"

	"github.com/polipastos/server/models"
)

func Test_ValidLogin(t *testing.T) {

	var err error
	var username, password string
	username = "validatelogin-usr"
	password = "validatelogin-pwd"

	var user models.User
	user.Username = username

	user.ID, err = RegisterUser(username, password)
	if err != nil {
		t.Fatalf("unable to create user %v", err)
	}

	err = db.Find(&user, user.ID)
	if err != nil {
		t.Fatal(err)
	}

	id, err := PairAuthentication(username, password)
	if err != nil {
		t.Fatal(err)
	}

	if id != user.ID {
		t.Fatalf("the returned id doesn't match with the user (%v / %v)", id, user.ID)
	}

}

func Test_UnregisteredUser(t *testing.T) {

	var err error
	var username, password string

	username = "unregistered-usr"
	password = "unregistered-pwd"

	_, err = PairAuthentication(username, password)

	if err != ErrNotFound {
		t.Fatal("error missed", err.Error())
	}
}

func Test_RegisteredUserInvalidPassword(t *testing.T) {

	var err error
	var username, password string
	username = "reg_user_inv_pw"
	password = "password_valid"

	var user models.User
	user.Username = username

	user.ID, err = RegisterUser(username, password)
	if err != nil {
		t.Fatal(err)
	}

	_, err = PairAuthentication(username, "thisisnot")
	if err == nil {
		t.Fatal("error missed")
	}

	if err != ErrInvalidPair {
		t.Fatalf("that's not the expected error %v", err)
	}
}
