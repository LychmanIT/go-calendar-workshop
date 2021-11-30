package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
)

type AuthorizeResponse struct {
	Token string `json:"token"`
}

var (
	svcURL = "http://" + net.JoinHostPort(os.Getenv("CALENDAR_CONTAINER_NAME"), os.Getenv("CALENDAR_HTTP_PORT"))
)

func TestHTTPStatuses(t *testing.T) {
	bearer := authorize()

	storeEventBody := `{"event":{"id":"test-1","title":"New Event","description":"qwewqe","time":"123","timezone":"America/Chicago","duration":2555,"notes":["Note","New Note"]}}`
	loginInvalidBody := `{"credentials":{"username": "gouser1", "password":"gopassword"}}`
	changeTimezoneBody := `{"timezone":"America/Chicago"}`
	updateEventBody := `{"event":{"title":"Updated Event","description":"asdasd","time":"321","timezone":"America/Toronto","duration":55,"notes":["qwe","asd"]}}`

	for _, testcase := range []struct {
		method, url, body, token string
		want                     int
	}{
		//	Status tests

		//Good request
		{"GET", svcURL + "/api/status", ``, "", http.StatusOK},
		//Bad request. Method POST is not allowed
		{"POST", svcURL + "/api/status", ``, "", http.StatusMethodNotAllowed},

		//	Login tests

		//Bad request. Bearer token is invalid
		{"POST", svcURL + "/login", loginInvalidBody, "", http.StatusForbidden},
		//Bad request. No user found
		{"POST", svcURL + "/login", loginInvalidBody, bearer, http.StatusForbidden},
		//Bad request. Method GET is not allowed
		{"GET", svcURL + "/login", loginInvalidBody, bearer, http.StatusMethodNotAllowed},

		//	ChangeTimezone tests

		//Good request
		{"POST", svcURL + "/change-timezone", changeTimezoneBody, bearer, http.StatusOK},
		//Bad request. Bearer token is invalid
		{"POST", svcURL + "/change-timezone", changeTimezoneBody, "", http.StatusForbidden},
		//Bad request. Method GET is not allowed
		{"GET", svcURL + "/change-timezone", changeTimezoneBody, bearer, http.StatusMethodNotAllowed},

		//	Store tests

		//Good request
		{"POST", svcURL + "/api/store", storeEventBody, bearer, http.StatusOK},
		//Bad request. Bearer token is invalid
		{"POST", svcURL + "/api/store", storeEventBody, "", http.StatusForbidden},
		//Bad request. Method PUT is not allowed
		{"PUT", svcURL + "/api/store", storeEventBody, bearer, http.StatusMethodNotAllowed},

		//	Show tests

		//Good request
		{"GET", svcURL + "/api/show/test-1", "", bearer, http.StatusOK},
		//Bad request. No event found
		{"GET", svcURL + "/api/show/test-2", "", bearer, http.StatusNotFound},
		//Bad request. Method PUT is not allowed
		{"PUT", svcURL + "/api/show/test-1", "", bearer, http.StatusMethodNotAllowed},
		//Bad request. Bearer token is invalid
		{"GET", svcURL + "/api/show/test-1", "", "", http.StatusForbidden},

		//	Index tests

		//Good request
		{"GET", svcURL + "/api/index", ``, bearer, http.StatusOK},
		//Good request. Filter usage
		{"GET", svcURL + "/api/index?description=qwewqe&time=123", ``, bearer, http.StatusOK},
		//Bad request. Method POST is not allowed
		{"POST", svcURL + "/api/index", ``, bearer, http.StatusMethodNotAllowed},
		//Bad request. Bearer token is not valid
		{"GET", svcURL + "/api/index", ``, "", http.StatusForbidden},

		//	Update tests

		//Bad request. No event found
		{"PUT", svcURL + "/api/update/test-2", updateEventBody, bearer, http.StatusNotFound},
		//Bad request. Method POST is not allowed
		{"POST", svcURL + "/api/update/test-1", updateEventBody, bearer, http.StatusMethodNotAllowed},
		//Bad request. Bearer token is not valid
		{"PUT", svcURL + "/api/update/test-1", updateEventBody, "", http.StatusForbidden},
		//Good request
		{"PUT", svcURL + "/api/update/test-1", updateEventBody, bearer, http.StatusOK},

		//	Delete tests

		//Bad request. No event found
		{"DELETE", svcURL + "/api/delete/test-2", "", bearer, http.StatusNotFound},
		//Bad request. Method POST is not allowed
		{"POST", svcURL + "/api/delete/test-1", "", bearer, http.StatusMethodNotAllowed},
		//Bad request. Bearer token is not valid
		{"DELETE", svcURL + "/api/delete/test-1", "", "", http.StatusForbidden},
		//Good request
		{"DELETE", svcURL + "/api/delete/test-1", "", bearer, http.StatusOK},

		//	Logout tests

		//Bad request. Method GET is not allowed
		{"GET", svcURL + "/logout", ``, bearer, http.StatusMethodNotAllowed},
		//Good request
		{"POST", svcURL + "/logout", ``, bearer, http.StatusOK},
	} {
		req, err := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		if err != nil {
			log.Println(err.Error())
		}
		req.Header.Add("Authorization", testcase.token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err.Error())
		}
		if testcase.want != resp.StatusCode {
			t.Errorf("%s %s %s: want %d have %d", testcase.method, testcase.url, testcase.body, testcase.want, resp.StatusCode)
		}
	}
}

func authorize() string {
	postBody, _ := json.Marshal(map[string]map[string]string{
		"credentials": {
			"username": "gouser",
			"password": "gopassword",
		},
	})

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(svcURL+"/login", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result AuthorizeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return "Bearer " + result.Token
}
