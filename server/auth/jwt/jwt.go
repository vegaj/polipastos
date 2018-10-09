package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

//Header is a jwt algorithm
type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}


//JWT json web token
type JWT struct {
	Header  Header                 `json:"header"`
	Payload map[string]interface{} `json:"payload"`
	//Signature []byte
}

//JSONTime to ensure that the time is parsed as a regular json object
type JSONTime time.Time

//MarshalJSON implements the json.Marshaler interface.
func (t JSONTime) MarshalJSON() ([]byte, error) {
	return t.MarshalJSON()
}

//UnmarshalJSON implements the json.Unmarshaler interface.
func (t JSONTime) UnmarshalJSON(data []byte) error {
	return t.UnmarshalJSON(data)
}

//NewJWT makes a new JSON web Token instance with all its fields
//created and set to blank. Creating the JWT with this method should
//prevent accessing to nil structures related panics.
func NewJWT() JWT {
	return JWT{
		Header:  Header{},
		Payload: make(Claims),
	}
}

//Encode generates the base64 urlEncoding form and its signature concatenadeted
//in the format <header>.<payload>.<signature> as described in jwt.io
//See https://wwww.jwt.io
func (jwt *JWT) Encode(secret []byte) string {

	jwt.Header.Algorithm = "HS256"
	jwt.Header.Type = "JWT"

	headJSON, _ := json.Marshal(&jwt.Header)
	payloadJSON, _ := json.Marshal(map[string]interface{}(jwt.Payload))

	baseHeader, basePayload :=
		base64.URLEncoding.EncodeToString(headJSON),
		base64.URLEncoding.EncodeToString(payloadJSON)

	message := baseHeader + "." + basePayload

	signature, err := Sign(message, secret)
	if err != nil {
		return ""
	}

	return message + "." + string(signature)
}

//FromString initializes the jwt with the information encoded in the input string
func (jwt *JWT) FromString(urlenc string, secret []byte) error {

	parts := strings.Split(urlenc, ".")

	if len(parts) != 3 {
		return errors.New("malformed token")
	}

	message := strings.Join(parts[:2], ".")

	if !Verify(message, parts[2], secret) {
		return errors.New("the message has been altered")
	}

	headerJSON := fromBase64String(parts[0])
	payloadJSON := fromBase64String(parts[1])

	//Header's fields are all strings so no additional cautions should be taken
	//into consideration.
	if err := json.Unmarshal(headerJSON, &jwt.Header); err != nil {
		return err
	}

	if jwt.Payload == nil {
		jwt.Payload = make(map[string]interface{})
	}

	//As some claims are int64 numbers and the default deserialization is as
	//float64, we let the numeric fields as json.Number, so the part that has
	//to access those claims must interpret it as desired.
	decoder := json.NewDecoder(bytes.NewBuffer(payloadJSON))
	decoder.UseNumber()
	var err = decoder.Decode(&jwt.Payload)

	//time.Time si serialized in json in a plain text format. So all the
	jwt.decodeTimeField(expirationTimeK)
	jwt.decodeTimeField(notBeforeK)
	jwt.decodeTimeField(issuedAtK)

	return err
}
