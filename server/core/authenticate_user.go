package core

import (
	"log"

	"github.com/gobuffalo/uuid"
	"github.com/polipastos/server/models"
)

//PairAuthentication checks if the given username and password are related and returns the uuid associated
//with that username
func PairAuthentication(username, password string) (uuid.UUID, error) {

	var user models.User
	var dummyID uuid.UUID

	if username == "" || password == "" {
		return dummyID, ErrInvalidRequest
	}

	e := db.Where("username = ?", username).First(&user)
	if e != nil {
		return dummyID, ErrNotFound
	}

	var digest models.PasswordDigest
	e = db.Where("owner_id = ?", user.ID).First(&digest)
	if e != nil {
		log.Println("unable to retrieve password digest:", e)
		return dummyID, ErrInternalError
	}

	if !PasswordMatch(digest.Digest, []byte(password)) {
		return dummyID, ErrInvalidPair
	}

	return user.ID, nil
}
