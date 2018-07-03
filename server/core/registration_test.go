package core

import (
	"testing"

	"github.com/gobuffalo/uuid"

	"github.com/polipastos/server/models"
)

var (
//db declared in core.go

)

func Test_ValidRegistration(t *testing.T) {

	var username, password string

	username = "Rodrigo"
	password = "Rodripass"

	id, err := RegisterUser(username, password)
	if err != nil {
		t.Fatal(err)
	}

	var user models.User
	err = db.Where("username = ?", username).First(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != id {
		t.Fatalf("the received id is not the expected %v / %v", user.ID, id)
	}
}

func Test_InvalidUsername(t *testing.T) {
	var username, password string

	password = "xxx"
	_, err := RegisterUser(username, password)
	if err == nil {
		t.Fatal("error missed > blank username")
	}

	username = "rock"
	password = ""
	id, err := RegisterUser(username, password)
	if err == nil {
		t.Errorf("error missed on blank password: id %v", id)
	}
}

func Test_DuplicatedRegistration(t *testing.T) {
	var username, password string

	username = "duplicatedreg"
	password = "duplicatedreg"

	_, err := RegisterUser(username, password)
	if err != nil {
		t.Fatal(err)
	}

	_, err = RegisterUser(username, password)
	if err == nil {
		t.Fatal("duplicated user registration missed")
	}
	//Raised error: mysql create: Error 1062: Duplicate entry 'duplicatedreg' for key 'username'
}

func Test_ValidRegistrationMatchesPassword(t *testing.T) {
	var username, password string

	username = "validregist"
	password = "validpassw"

	id, err := RegisterUser(username, password)
	if err != nil {
		t.Fatal(err)
	}

	var pwdig models.PasswordDigest
	err = db.Where("owner_id = ?", id).First(&pwdig)
	if err != nil {
		t.Fatalf("entry expected: %v", err)
	}

	if !PasswordMatch(pwdig.Digest, []byte(password)) {
		t.Fatalf("the password cannot ve berified (stored one) %v...", pwdig.Digest[:8])
	}
}

func Test_CannotAccessOtherUsersOtherPassword(t *testing.T) {
	var unam1, unam2, passw1, passw2 string

	unam1 = "pwcrossaccess1"
	unam2 = "pwcrossaccess2"
	passw1 = "pw1"
	passw2 = "pw2"

	var err error
	var id, id1, id2 uuid.UUID

	id1, err = RegisterUser(unam1, passw1)
	if err != nil {
		t.Fatal(err)
	}

	id2, err = RegisterUser(unam2, passw2)
	if err != nil {
		t.Fatal(err)
	}

	if id1 == id2 {
		t.Fatalf("different users with same id %v | %v", id1, id2)
	}

	id, err = PairAuthentication(unam1, passw2)
	if err == nil {
		t.Fatal("the pair authentication should'd fail")
	} else if ErrInvalidPair != err {
		t.Fatalf("%v error was expected, but found %v", ErrInvalidPair, err)
	}

	if id == id1 || id == id2 {
		t.Fatalf("a valid uuid was retourned (%v) u1 (%v) u2 (%v)", id, id1, id2)
	}

}
