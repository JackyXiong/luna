package luna

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

type Response struct {
	*http.Response
	content []byte
}

// return response as []byte
func (r Response) Content() (content []byte, err error) {
	if r.content != nil {
		return r.content, nil
	}
	var reader io.ReadCloser
	switch r.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
	case "deflate":
		reader, err = zlib.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = r.Body
	}
	defer reader.Close()

	if content, err = ioutil.ReadAll(reader); err != nil {
		return nil, err
	}
	r.content = content
	return content, nil
}

//return response as text
func (r Response) Text() (text string, err error) {
	content, err := r.Content()
	if err != nil {
		return "", err
	}
	text = string(content)
	return
}

//return response as json
func (r Response) Json() (json *simplejson.Json, err error) {
	content, err := r.Content()
	if err != nil {
		return "", err
	}
	return simplejson.NewJson(content)
}
