package fetch

import (
	"encoding/json"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
	"topheruk.com/encode-json/domain/googleapi"
)

func TestFetch(t *testing.T) {
	var url = `https://www.googleapis.com/books/v1/volumes?q=isbn:9780358380243`

	var api googleapi.Books
	err := Fetch(url, func(resp *http.Response) error {
		return json.NewDecoder(resp.Body).Decode(&api)
	}, DefaultOptions)
	assert.Assert(t, err)

	assert.Assert(t, cmp.Equal(api.Items[0].VolumeInfo.Title, "The Two Towers"))
}
