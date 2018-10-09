package core

import (
	"os"
	"testing"

	"github.com/polipastos/server/test"
)

func TestMain(m *testing.M) {
	var err error
	db, env, err = test.PrepareTests()

	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
