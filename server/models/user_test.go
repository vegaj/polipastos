package models

import (
	"os"
	"testing"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/polipastos/server/test"
)

var (
	db  *pop.Connection
	env test.Env
)

const (
	targetdb = "testing"
)

func TestMain(m *testing.M) {
	var err error
	db, env, err = test.PrepareTests()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func Test_CreateUser(t *testing.T) {

	var u User
	var err error

	u.Username = "Cabashero"

	errors, err := u.ValidateCreate(db)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if errors.HasAny() {
		t.Fatalf("%v", errors.Error())
	}

	err = db.Create(&u)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CheckUserValidation(t *testing.T) {
	var u User

	//Not set
	u.Username = ""

	errs, err := db.ValidateAndCreate(&u)

	if !errs.HasAny() {
		t.Fatal("Expected an error because username has been left blank")
	}

	if err != nil {
		t.Fatal("The validation error was expected to be present in errors, not here")
	}
}

func Test_TimestampIgnored(t *testing.T) {
	var u, u2 User

	dur, _ := time.ParseDuration("1h")

	u.Username = "Timestampito"
	u.CreatedAt = time.Now().Add(dur)
	u.UpdatedAt = time.Now().Add(dur)

	errs, err := db.ValidateAndCreate(&u)
	t.Logf("errors <%v>. err <%v>", errs, err)

	db.Find(&u2, u.ID)
	if !u.CreatedAt.After(u2.CreatedAt) {
		t.Errorf("the created_at field should'd been ignored and updated by pop")
	}

	if !u.UpdatedAt.After(u2.UpdatedAt) {
		t.Errorf("the updated_at field should'd been ignored and updated by pop")
	}
}

func Test_UpdateValidation(t *testing.T) {
	var pepe, pepito User

	pepe.Username = "pepe"
	pepito.Username = "pepe"

	errs, err := db.ValidateAndCreate(&pepe)
	if errs.HasAny() || err != nil {
		t.Fatal("no errors expected")
	}

	errs, err = db.ValidateAndCreate(&pepito)
	if !errs.HasAny() && err == nil {
		t.Fatal("Unique name collision undetected")
	}

	pepe.Username = "pepo"
	errs, err = db.ValidateAndUpdate(&pepe)
	if errs.HasAny() || err != nil {
		t.Fatal("no errors expected")
	}

	errs, err = db.ValidateAndCreate(&pepito)
	if errs.HasAny() || err != nil {
		t.Fatal("no errors expected")
	}
}
