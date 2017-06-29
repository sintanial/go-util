package httputil

import (
	"net/http"
	"strings"
	"compress/gzip"
	"io"
	"github.com/ansel1/merry"
	"io/ioutil"
)

func GzipWrite(w http.ResponseWriter, r *http.Request, minbytes int, data []byte) error {
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") && len(data) > minbytes {
		w.Header().Set("Content-Encoding", "gzip")
		gzipper := gzip.NewWriter(w)
		_, err := gzipper.Write(data)
		gzipper.Close()

		return err
	} else {
		_, err := w.Write(data)
		return err
	}
}

func GzipReader(r *http.Request) (io.Reader, error) {
	body := r.Body
	if r.Header.Get("Content-Encoding") == "gzip" {
		var err error
		body, err = gzip.NewReader(r.Body)
		if err != nil {
			return nil, merry.Wrap(err)
		}
	}

	return body, nil
}

func ReadAll(r *http.Request) ([]byte, error) {
	body, err := GzipReader(r)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(body)
}

func ReadAllString(r *http.Request) (string, error) {
	data, err := ReadAll(r)
	return string(data), err
}

