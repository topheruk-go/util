package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func htmlFormattedTimeDuration(t time.Time) (start, end string) {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d), fmt.Sprintf("%d-%02d-%02d", y+10, m, d)
}

func parseMultiPartForm(rw http.ResponseWriter, r *http.Request, lf *LoanFormDto) error {
	if err := r.ParseMultipartForm(1e6); err != nil {
		return err
	}
	f, _, err := r.FormFile("loan-form")
	if err != nil {
		return err
	}
	defer f.Close()
	lf.StudentID, lf.StartDate, lf.EndDate, err = parseForm(rw, r)
	if err != nil {
		return err
	}
	return uploadFile(f, lf)
}

func postLoanForm(url, contentType string, lf *LoanFormDto) (err error) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(lf)
	_, err = http.Post(url, contentType, &buf)
	return
}

func parseForm(rw http.ResponseWriter, r *http.Request) (userID string, startDate, endDate *time.Time, err error) {
	layout := "2006-01-02"
	var ts time.Time
	if ts, err = time.Parse(layout, r.PostFormValue("start-date")); err != nil {
		return "", nil, nil, fmt.Errorf("error parsing start-date field: %w", err)
	}
	var te time.Time
	if te, err = time.Parse(layout, r.PostFormValue("end-date")); err != nil {
		return "", nil, nil, fmt.Errorf("error parsing end-date field: %w", err)
	}
	return r.PostFormValue("user-id"), &ts, &te, nil
}

// currently for every document I have to write the content to a new file
// it would be more effective to only do this when fetching data on a single upload
// this would require a complete rewrite though
// the files are currently not large enough where this is a concern
func uploadFile(f multipart.File, lf *LoanFormDto) error {
	tf, err := os.CreateTemp(lf.TmpPath, "*.pdf")
	if err != nil {
		return err
	}
	defer tf.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if _, err = tf.Write(b); err != nil {
		return err
	}
	lf.TmpPath = tf.Name()
	return nil
}

func parseURL(r *http.Request, path string) string {
	// FIXME:
	scheme := strings.ToLower(strings.Split(r.Proto, "/")[0])
	return fmt.Sprintf("%s://%s:%s", scheme, r.Host, path)
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
