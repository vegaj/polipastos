package auth

import (
	"testing"
	"time"
)

/*
	The json serialization of a map[string] interface{} makes the
	decoder of the json assert which types each field is.
	By default all numerical numbers are parsed as float64 so
	if you want to send a int32 integer, you have to parse the
	received float64 as int64, then cast it to int32 before using it.
*/
func Test_JWT_NumericClaims(t *testing.T) {

	var inToken = NewJWT()

	t1, t2 := time.Now(), time.Now()

	inToken.SetAudience([]string{"fido-issuer", "eminem-user", "small-city town"})
	inToken.SetNotBefore(t1)
	inToken.SetIssuedAt(t1)
	inToken.SetExpirationTime(t2)

	secret := []byte("very-very-secret-my-friend")
	decoded := NewJWT()

	decoded.FromString(inToken.Encode(secret), secret)

	if decoded.ExpirationTime().Before(t2) {
		t.Log("Before")
	}

	if decoded.ExpirationTime().After(t2) {
		t.Log("After")
	}

	if decoded.ExpirationTime().Equal(t2) {
		t.Errorf("Expiration token not received succesfully")
	}

	if decoded.NotBefore() != t1 {
		t.Errorf("NotBefore not received succesfully")
	}

	if decoded.IssuedAt() != t1 {
		t.Errorf("NotBefore not received succesfully")
	}

	if len(decoded.Audience()) != 3 {
		t.Errorf("the audience has not the expected fields")
	}
}

func Test_JWT_NumericalClaimsNotPresent(t *testing.T) {

	var token = NewJWT()
	if v := token.ExpirationTime(); !v.IsZero() {
		t.Fatalf("optional field (expiration) should have returned default time. %v instead", v)
	}

	if v := token.NotBefore(); !v.IsZero() {
		t.Fatalf("optional field (not before) shoul'd have returned 0. %d instead", v)
	}

	if v := token.IssuedAt(); !v.IsZero() {
		t.Fatalf("optional field (issuedAt) shoul'd have returned 0. %d instead", v)
	}

}

func Test_JWT_Validation(t *testing.T) {

	secret := []byte("the secret here")

	ref := time.Now()

	minute, _ := time.ParseDuration("1m")
	twomin, _ := time.ParseDuration("2m")

	t1 := ref.Add(minute)
	t2 := ref.Add(twomin)

	b := NewJWT()
	token := NewJWT()
	token.SetExpirationTime(t2)
	token.SetNotBefore(t1)
	b.FromString(token.Encode(secret), secret)

	if e := b.Validate(t2); e != ErrNotBefore {
		t.Errorf("expected err not before <%v>", e)
	}

	b = NewJWT()
	//token = NewJWT()

	token.SetExpirationTime(t1)
	token.SetNotBefore(ref)
	b.FromString(token.Encode(secret), secret)

	if e := b.Validate(t2); e != ErrExpiredToken {
		t.Errorf("expected expired token error <%v>", e)
	}

	b = NewJWT()
	token.SetExpirationTime(t2)
	token.SetNotBefore(ref)
	b.FromString(token.Encode(secret), secret)
	if e := b.Validate(t1); e != nil {
		t.Errorf("the token has been rejected and was valid. error <%v>", e)
	}
}
