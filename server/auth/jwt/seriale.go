package auth

import (
	"encoding/json"
	"strings"
	"time"
)

func obtainTimeFromText(text []byte) time.Time {

	var t time.Time
	t.UnmarshalText(text)
	return t

}

//MarshalJSON implements json.Marshaler
func (c Claims) MarshalJSON() ([]byte, error) {
	var objects []string
	var data []byte
	var err error
	for k, v := range c {

		if k == expirationTimeK || k == notBeforeK || k == issuedAtK {
			data, err = v.(time.Time).MarshalJSON()
		} else {
			data, err = json.Marshal(v)
		}

		if err != nil {
			panic(err)
		}

		objects = append(objects, string(data))
	}

	return []byte(strings.Join(objects, ",")), err
}

func (jwt *JWT) decodeTimeField(fieldKey string) {

	//check if the field is present
	if field, ok := jwt.Payload[fieldKey]; ok {

		if asTime, ok := field.(time.Time); ok {
			//the field has been set as time.Time
			jwt.Payload[fieldKey] = asTime
		} else if asString, ok := field.(string); ok {
			//the field is a string
			jwt.Payload[fieldKey] = obtainTimeFromText([]byte(asString))
		} else {
			//The time format is not recognized.
			//We do nothing
		}
	}
}
