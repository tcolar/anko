// +build !appengine

// Package http implements http interface for anko script.
package http

import (
	"errors"
	"github.com/mattn/anko/vm"
	h "net/http"
	"reflect"
)

type Client struct {
	c *h.Client
}

func (c *Client) Get(args ...reflect.Value) (reflect.Value, error) {
	if len(args) < 1 {
		return vm.NilValue, errors.New("Missing arguments")
	}
	if len(args) > 1 {
		return vm.NilValue, errors.New("Too many arguments")
	}
	if args[0].Kind() != reflect.String {
		return vm.NilValue, errors.New("Argument should be string")
	}
	res, err := h.Get(args[0].String())
	return reflect.ValueOf(res), err
}

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("DefaultClient", reflect.ValueOf(h.DefaultClient))
	m.Define("NewServeMux", reflect.ValueOf(h.NewServeMux))
	m.Define("Handle", reflect.ValueOf(h.Handle))
	m.Define("HandleFunc", reflect.ValueOf(h.HandleFunc))
	m.Define("ListenAndServe", reflect.ValueOf(h.ListenAndServe))
	return m
}
