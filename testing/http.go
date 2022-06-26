package testing

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gotest.tools/v3/assert"

	ptth "github.com/topheruk-go/util/http"
)

func HttpService(t *testing.T, srv *httptest.Server, tt []ptth.TestCase) {
	for _, tc := range tt {
		if tc.ContentType == "" {
			tc.ContentType = "application/json"
		}
		if tc.Status == 0 {
			tc.Status = http.StatusOK
		}
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest(tc.Method, srv.URL+tc.PathName, strings.NewReader(tc.Content))
			assert.Equal(t, err, nil)

			req.Header.Add("Content-Type", tc.ContentType)

			res, err := srv.Client().Do(req)
			assert.Equal(t, err, nil)

			assert.Equal(t, res.StatusCode, tc.Status)
		})
	}
}
