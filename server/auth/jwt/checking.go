package auth

//Package auth/jwt will be removed in the future. JOSE should be used instead.

import (
	"errors"
	"time"
)

var (
	//ErrExpiredToken the token must be discarded
	ErrExpiredToken = errors.New("expired token")
	//ErrNotBefore the token is not valid yet
	ErrNotBefore = errors.New("token not valid yet")
	//NilT time
	NilT = time.Time{}
)

//Validate the token and returns nil on success
//The reference time (can be time.Now()) is passed through t
func (jwt JWT) Validate(t time.Time) error {

	if err := jwt.checkExpirationTime(t); err != nil {
		return err
	}

	if err := jwt.checkNotBefore(t); err != nil {
		return err
	}

	return nil
}

//checkExpirationTime respect to the time it arrived.
func (jwt JWT) checkExpirationTime(receivedAt time.Time) error {

	//Optional field: 0 means not present or invalid, so ignored
	if jwt.ExpirationTime().IsZero() {
		return nil
	}

	if jwt.ExpirationTime().After(receivedAt) {
		return ErrExpiredToken
	}
	return nil
}

func (jwt JWT) checkNotBefore(receivedAt time.Time) error {

	/*Optional field so not taken into consideration
	if jwt.NotBefore().isZero() {
		return nil
	}
	*/

	if time.Now().Before(jwt.NotBefore()) {
		return ErrNotBefore
	}
	return nil
}
