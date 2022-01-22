package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/topheruk/go/example/app/sql-01/model/v1"
)

func timeDuration(t time.Time) (start, end string) {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d), fmt.Sprintf("%d-%02d-%02d", y+10, m, d)
}

func parseMultiPartForm(rw http.ResponseWriter, r *http.Request, lf *model.LoanForm) error {
	if err := r.ParseMultipartForm(1e6); err != nil {
		return fmt.Errorf("error parsing form: %w", err)
	}

	f, _, err := r.FormFile("loan-form")
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}
	defer f.Close()

	lf.StudentID, lf.StartDate, lf.EndDate, err = parseForm(rw, r)
	if err != nil {
		return err
	}

	if err := uploadFile(f, lf); err != nil {
		return err
	}

	return nil
}

func postLoanForm(url, contentType string, lf *model.LoanForm) error {
	b, err := json.Marshal(lf)
	if err != nil {
		return fmt.Errorf("error marshaling loan-form data:%w", err)
	}
	buf := bytes.NewBuffer(b)

	_, err = http.Post(url, contentType, buf)
	if err != nil {
		return fmt.Errorf("error making post request:%w", err)
	}

	return nil
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

func uploadFile(f multipart.File, lf *model.LoanForm) error {
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
