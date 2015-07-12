package main

import (
	"github.com/manyminds/api2go/jsonapi"
	"net/http"
	"strconv"
)

type JsonApiBinding struct{}

func (JsonApiBinding) Name() string {
	return "jsonapi"
}

func (JsonApiBinding) Bind(req *http.Request, obj interface{}) error {
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	formData := make([]byte, length)
	_, err = req.Body.Read(formData)
	if err != nil {
		if err.Error() != "EOF" {
			return err
		}
	}

	err = jsonapi.UnmarshalFromJSON(formData, obj)
	if err != nil {
		return err
	}

	return nil
}
