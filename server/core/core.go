package core

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/gobuffalo/pop"
	"github.com/polipastos/server/test"
)

var (
	db  *pop.Connection
	env test.Env

	//ErrAlreadyInUse shows a pk conflict
	ErrAlreadyInUse = errors.New("identifier already in use")
	//ErrInternalError shows that the error is on the server side
	ErrInternalError = errors.New("internal server error")
	//ErrInvalidRequest - The arguments for this function are invalid
	ErrInvalidRequest = errors.New("invalid request")
	//ErrNotFound - The requested resource could not be found with the given information
	ErrNotFound = errors.New("resource not found")
	//ErrInvalidPair - The pair authentication failed.
	ErrInvalidPair = errors.New("invalid credentials")
)

//Init the core package with the persistence database connection
func Init(popconn *pop.Connection) {
	if popconn == nil {
		panic("Please, provide a valid database connection <nil>")
	}

	if db == nil {
		db = popconn
	}

}

//PasswordMatch checks if the provided password generates the digested one
func PasswordMatch(digest, password []byte) bool {
	return bcrypt.CompareHashAndPassword(digest, password) == nil
}
