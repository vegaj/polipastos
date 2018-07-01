package models

import (
	"os"

	"github.com/gobuffalo/pop"
)

type testEnv map[string]string

func prepareTests() (*pop.Connection, testEnv, error) {

	var err error
	var db *pop.Connection
	var migrator pop.FileMigrator

	env := make(map[string]string)
	env["migrations"] = os.Getenv("migrations_path")
	env["targetdb"] = os.Getenv("database")
	env["pwd"], err = os.Getwd()
	//env["new_key"] = "" for every new env to be taken into consideration

	if err != nil {
		return nil, env, err
	}

	db, err = pop.Connect(env["targetdb"])

	migrator, err = pop.NewFileMigrator(env["migrations"], db)
	if err != nil {
		return nil, env, err
	}

	err = migrator.Reset()
	if err != nil {
		return nil, env, err
	}

	err = migrator.Status()
	if err != nil {
		return nil, env, err
	}

	return db, env, err
}
