package luna

import (
	"bytes"
	"encoding/json"
	"errors"
	// "fmt"
	"io"
	"mime/multipart"
	"os"
)

func newMultipartBody(reqOpt *ReqOptions) (body io.Reader, contentType string, err error) {
	contentType = "application/form-data"
	if len(reqOpt.Files) == 0 {
		return nil, "", errors.New("files unexists")
	}
	bodyBuf := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, file := range reqOpt.Files {
		fPart, err := bodyWriter.CreateFormFile(file.Name, file.Path)
		if err != nil {
			return nil, contentType, err
		}
		file, err := os.Open(file.Path)
		defer file.Close()

		if err != nil {
			return nil, contentType, err
		}
		_, err = io.Copy(fPart, file)
		if err != nil {
			return nil, contentType, err
		}
	}
	contentType = bodyWriter.FormDataContentType()

	// Close
	err = bodyWriter.Close()
	if err != nil {
		return nil, "", err
	}

	body = bodyBuf
	return
}

func newJsonBody(reqOpt *ReqOptions) (body io.Reader, contentType string, err error) {
	if reqOpt.Json == nil {
		return nil, "", errors.New("reqOpt has no Json field")
	}
	contentType = "application/json"
	b, err := json.Marshal(reqOpt.Json)
	if err != nil {
		return nil, "", err
	}
	body = bytes.NewReader(b)
	// fmt.Println(body)
	return
}
