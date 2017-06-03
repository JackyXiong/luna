package luna

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

// Response wrap origin http Response
type Response struct {
	Resp    *http.Response
	content []byte
	History []*http.Response // redirect history
}

// Content return response as []byte
func (r Response) Content() (content []byte, err error) {
	if r.content != nil {
		return r.content, nil
	}
	var reader io.ReadCloser
	switch r.Resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(r.Resp.Body)
		if err != nil {
			return nil, err
		}
	case "deflate":
		reader, err = zlib.NewReader(r.Resp.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = r.Resp.Body
	}
	defer reader.Close()

	if content, err = ioutil.ReadAll(reader); err != nil {
		return nil, err
	}
	r.content = content
	return content, nil
}

// Text return response as text
func (r Response) Text() (text string, err error) {
	content, err := r.Content()
	if err != nil {
		return "", err
	}
	text = string(content)
	return
}

// JSON return response as json
func (r Response) JSON() (json *simplejson.Json, err error) {
	content, err := r.Content()
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(content)
}
