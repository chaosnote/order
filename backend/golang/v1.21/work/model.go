package main

import (
	"idv/chris/component"
	"idv/chris/errs"
)

const (
	OK errs.HttpResponseCode = "OK"
)

type HttpResponse struct {
	Code    string
	Content map[string]interface{}
}

//-----------------------------------------------

var comps component.Store
