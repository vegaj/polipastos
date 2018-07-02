package core

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/gobuffalo/pop"
)

var (
	db *pop.Connection

	//ErrAlreadyInUse shows a pk conflict
	ErrAlreadyInUse = errors.New("identifier already in use")
	//ErrInternalError shows that the error is on the server side
	ErrInternalError = errors.New("internal server error")
	//ErrInvalidRequest - The arguments for this function are invalid
	ErrInvalidRequest = errors.New("invalid request")
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
