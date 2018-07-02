package core

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/gobuffalo/pop"

	"github.com/gobuffalo/uuid"
	"github.com/polipastos/server/models"
)

//RegisterUser will check if the username is taken, digest the password and
//create a new user in the database that could be found by teh uuid returned
func RegisterUser(username, password string) (uuid.UUID, error) {

	var u models.User

	if username == "" || password == "" {
		return uuid.UUID{}, ErrInvalidRequest
	}

	u.Username = username
	err := db.Transaction(func(c *pop.Connection) error {

		errs, e := c.ValidateAndCreate(&u)
		if errs.HasAny() || e != nil {
			log.Println("Unable to create user:", errs.Error(), "\n", e.Error())
			return ErrAlreadyInUse
		}

		var digest []byte
		digest, e = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if e != nil {
			return e
		}

		var pwdig models.PasswordDigest
		pwdig.Digest = digest
		pwdig.OwnerID = u.ID

		errs, e = c.ValidateAndCreate(&pwdig)
		if errs.HasAny() || e != nil {
			log.Println("Password association problem", errs.Error(), "\n", e.Error())
			return ErrInternalError
		}
		return nil
	})

	return u.ID, err
}
