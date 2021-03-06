// Package path implements path interface for anko script.
package path

import (
	"github.com/mattn/anko/vm"
	pkg "path"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("Base", reflect.ValueOf(pkg.Base))
	m.Define("Clean", reflect.ValueOf(pkg.Clean))
	m.Define("Dir", reflect.ValueOf(pkg.Dir))
	m.Define("ErrBadPattern", reflect.ValueOf(pkg.ErrBadPattern))
	m.Define("Ext", reflect.ValueOf(pkg.Ext))
	m.Define("IsAbs", reflect.ValueOf(pkg.IsAbs))
	m.Define("Join", reflect.ValueOf(pkg.Join))
	m.Define("Match", reflect.ValueOf(pkg.Match))
	m.Define("Split", reflect.ValueOf(pkg.Split))
	return m
}
