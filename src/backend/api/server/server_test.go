package server

import (
	"bytes"
	internal "dog/internal/app/common"
	"dog/internal/app/dig/dtos"
	"dog/internal/app/dig/queries"
	"encoding/json"
	"flag"
	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type TestDouble struct {
	Server *Server
}

var Mock *TestDouble

func TestMain(m *testing.M) {
	Mock = new(TestDouble)
	Mock.Server = NewServer("../../configs/app_config.yml")
	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestDig(t *testing.T) {
	type test struct {
		query    queries.DigQuery
		response *dtos.DigResponse
		code     int
		message  string
		assert   func() bool
	}

	tests := []test{
		{
			query:    queries.DigQuery{Domain: "google.com", Type: dns.TypeA},
			code:     http.StatusOK,
			message:  "get dig information for a valid query was successful",
			response: &dtos.DigResponse{},
		},
		{
			query:    queries.DigQuery{Domain: "google.com", Type: dns.TypeAAAA},
			code:     http.StatusOK,
			message:  "get dig information for a valid TXT query was successful",
			response: &dtos.DigResponse{},
		},
	}
	tests[0].assert = func() bool {
		return len(tests[0].response.Records) > 0
	}
	tests[1].assert = func() bool {
		return len(tests[0].response.Records) > 0
	}

	for _, tc := range tests {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(tc.query)
		req, _ := http.NewRequest("POST", "/api/v1/dig/", b)
		w := httptest.NewRecorder()
		Mock.Server.Router.ServeHTTP(w, req)

		json.NewDecoder(w.Body).Decode(&tc.response)

		assert.Equal(t, tc.code, w.Code)
		assert.True(t, tc.assert())

	}

	t.Run("get dig information for a valid query", func(t *testing.T) {

		query := queries.DigQuery{Type: dns.TypeA}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(query)
		req, _ := http.NewRequest("POST", "/api/v1/dig/", b)
		w := httptest.NewRecorder()
		Mock.Server.Router.ServeHTTP(w, req)

		response := internal.CommonError{}
		json.NewDecoder(w.Body).Decode(&response)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}
