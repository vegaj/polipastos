package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/polipastos/server/core"
	"github.com/polipastos/server/test"
)

func TestMain(m *testing.M) {
	var err error
	db, env, err = test.PrepareTests()

	if err != nil {
		panic(err)
	}

	core.Init(db)

	os.Exit(m.Run())
}

func Test_RegisterValid(t *testing.T) {

	r, err := http.NewRequest(http.MethodPost, "/api/register",
		bytes.NewBufferString(url.Values{
			"username": {"pepe"},
			"password": {"pepepw"},
		}.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := httptest.NewRecorder()
	RegisterUserPost(tr, r)

	res := tr.Result()
	body, _ := ioutil.ReadAll(res.Body)

	var m url.Values
	err = json.Unmarshal(body, &m)
	if err != nil {
		t.Fatal(err)
	}

	id := m.Get("uuid")

	if id == "" {
		t.Fatal("missing key id")
	}

	t.Log(id)
}

func Test_WithApplicationUrlencoded(t *testing.T) {
	form := make(url.Values)

	form.Add("username", "fido-json-enc-header")
	form.Add("password", "fidopass")

	r, err := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBufferString(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	RegisterUserPost(w, r)

	m := make(url.Values)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &m)
	if err != nil {
		t.Fatalf("error deserializing the response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected StausOK (%d) Found %d", http.StatusOK, resp.StatusCode)
	}

	id := m.Get("uuid")
	ok, err := strconv.ParseBool(m.Get("ok"))

	if err != nil {
		t.Fatalf("the response has a flag that could not be parsed to the right type %v", err)
	}

	if !ok {
		t.Fatal("the response contains an error flag")
	}

	if id == "" {
		t.Fatal("the uuid field is not present")
	}

}

func Test_RegisterWithNoUsername(t *testing.T) {

	form := make(url.Values)
	form.Add("password", "fido")

	r, err := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBufferString(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	RegisterUserPost(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status code %d, found %d", http.StatusBadRequest, resp.StatusCode)
	}

	m := make(url.Values)
	err = json.Unmarshal(body, &m)
	if err != nil {
		t.Fatalf("error deserializing the response: %v", err)
	}

	message := m.Get("error")
	if message == "" {
		t.Fatal("the error was not reported in the body")
	} else if message != ErrInvalidRequest.Error() {
		t.Fatal("the reported message is not the expected one")
	}
}
