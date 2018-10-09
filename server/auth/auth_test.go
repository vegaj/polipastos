package auth

import (
	"os"
	"testing"

	"github.com/gobuffalo/pop"

	"github.com/polipastos/server/test"
)

var (
	db  *pop.Connection
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
