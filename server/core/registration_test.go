package core

import (
	"os"
	"testing"

	"github.com/polipastos/server/models"
	"github.com/polipastos/server/test"
)

var (
	//db declared in core.go
	env test.Env
)

func TestMain(m *testing.M) {
	var err error
	db, env, err = test.PrepareTests()

	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

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
