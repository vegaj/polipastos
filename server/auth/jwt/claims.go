package auth

import (
	"time"
)

var (
	//Registered claim names
	issuerK         = "iss"
	subjectK        = "sub"
	audienceK       = "aud"
	expirationTimeK = "exp"
	notBeforeK      = "nbf"
	issuedAtK       = "iat"
	jwtIDK          = "jti"

	//public claim names see https://www.iana.org/assignments/jwt/jwt.xhtml
)

//Claims are taken from the specification of JWT accessible here:
//https://tools.ietf.org/html/rfc7519#section-4.1
//the documentation here, thus, is a piece of that document.
//It's encouraged to access the registered claims with the method provided
//in this package and using
type Claims map[string]interface{}

//Issuer claim identifies the principal that issued the JWT.
func (jwt JWT) Issuer() string {
	if jwt.Payload == nil {
		return ""
	}
	return jwt.Payload[issuerK].(string)
}

//SetIssuer for this token with its provided string or uri.
func (jwt *JWT) SetIssuer(issuer string) {
	jwt.Payload[issuerK] = issuer
}

//Subject claim identifies the principal that is the subject of the JWT.
func (jwt JWT) Subject() string {
	if jwt.Payload == nil {
		return ""
	}
	return jwt.Payload[subjectK].(string)
}

//SetSubject for this token with its provided string or uri.
func (jwt *JWT) SetSubject(subject string) {
	jwt.Payload[issuerK] = subject
}

//Audience is the number of entities that are supposed to accept this jwt.
func (jwt JWT) Audience() []string {
	if jwt.Payload == nil {
		return nil
	}

	slice := jwt.Payload[audienceK].([]interface{})

	strlist := make([]string, len(slice))
	for i, v := range slice {
		strlist[i] = v.(string)
	}

	return strlist
}

//SetAudience for this tokens with the stringOrUris provided.
//Smashes any possible previous value
func (jwt *JWT) SetAudience(aud []string) {
	jwt.Payload[audienceK] = aud
}

//ExpirationTime claim identifies the expiration time on
//or after which the JWT MUST NOT be accepted for processing.
//If the claim is not correct or is not present, returns 0.
func (jwt JWT) ExpirationTime() time.Time {
	if jwt.Payload == nil {
		return time.Time{}
	}

	if field, ok := jwt.Payload[expirationTimeK]; ok {
		if time, ok := field.(time.Time); ok {
			return time
		}
	}

	return time.Time{}
}

//SetExpirationTime for this token
func (jwt *JWT) SetExpirationTime(exp time.Time) {
	jwt.Payload[expirationTimeK] = exp
}

//NotBefore claim identifies the time before which the JWT MUST NOT be accepted for processing
//If the field is not present or is incorrect, 0 will be returned
func (jwt JWT) NotBefore() time.Time {
	if jwt.Payload == nil {
		return NilT
	}

	if field, ok := jwt.Payload[notBeforeK]; ok {
		if time, ok := field.(time.Time); ok {
			return time
		}
	}

	return NilT
}

//SetNotBefore time for this token
func (jwt *JWT) SetNotBefore(nbf time.Time) {
	jwt.Payload[notBeforeK] = nbf
}

//IssuedAt claim identifies the time at wich the JWT was issued.
func (jwt JWT) IssuedAt() time.Time {
	if jwt.Payload == nil {
		return NilT
	}

	if field, ok := jwt.Payload[issuedAtK]; ok {
		if time, ok := field.(time.Time); ok {
			return time
		}
	}

	return NilT
}

//SetIssuedAt time for this token
func (jwt *JWT) SetIssuedAt(iat time.Time) {
	jwt.Payload[issuedAtK] = iat
}

//ID refers to the claim jti wich is the JWT ID.
func (jwt JWT) ID() string {
	return jwt.Payload[jwtIDK].(string)
}

//SetID for this token
func (jwt *JWT) SetID(jti string) {
	jwt.Payload[jwtIDK] = jti
}

//DelIssuer removes this entry
func (jwt *JWT) DelIssuer() {
	delete(jwt.Payload, issuerK)
}

//DelSubject removes this entry
func (jwt *JWT) DelSubject() {
	delete(jwt.Payload, subjectK)
}

//DelAudience removes all the audiences for this token
func (jwt *JWT) DelAudience() {
	delete(jwt.Payload, audienceK)
}

//DelExpirationTime removes this entry
func (jwt *JWT) DelExpirationTime() {
	delete(jwt.Payload, expirationTimeK)
}

//DelNotBefore removes this entry
func (jwt *JWT) DelNotBefore() {
	delete(jwt.Payload, notBeforeK)
}

//DelIssuedAt removes this entry
func (jwt *JWT) DelIssuedAt() {
	delete(jwt.Payload, issuedAtK)
}

//DelID removes the entry that matches the claim jti
func (jwt *JWT) DelID() {
	delete(jwt.Payload, jwtIDK)
}
