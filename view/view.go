// Copyright 2017 Gerasimos Maropoulos, ΓΜ. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package view

import (
	"io"
	"path/filepath"

	"github.com/denniselite/iris-fixed/core/errors"
)

// View is responsible to
// load the correct templates
// for each of the registered view engines.
type View struct {
	engines []Engine
}

// Register loads all the view engines' template files or embedded assets.
func (v *View) Register(e Engine) error {
	v.engines = append(v.engines, e)
	return nil
}

// Find receives a filename, gets its extension and returns the view engine responsible for that file extension
func (v *View) Find(filename string) Engine {
	extension := filepath.Ext(filename)
	// Read-Only no locks needed, at serve/runtime-time the library is not supposed to add new view engines
	for i, n := 0, len(v.engines); i < n; i++ {
		e := v.engines[i]
		if e.Ext() == extension {
			return e
		}
	}
	return nil
}

// Len returns the length of view engines registered so far.
func (v *View) Len() int {
	return len(v.engines)
}

var (
	errNoViewEngineForExt = errors.New("no view engine found for '%s'")
)

// ExecuteWriter calls the correct view Engine's ExecuteWriter func
func (v *View) ExecuteWriter(w io.Writer, filename string, layout string, bindingData interface{}) error {
	e := v.Find(filename)
	if e == nil {
		return errNoViewEngineForExt.Format(filepath.Ext(filename))
	}

	return e.ExecuteWriter(w, filename, layout, bindingData)
}

// AddFunc adds a function to all registered engines.
// Each template engine that supports functions has its own AddFunc too.
func (v *View) AddFunc(funcName string, funcBody interface{}) {
	for i, n := 0, len(v.engines); i < n; i++ {
		e := v.engines[i]
		if engineFuncer, ok := e.(EngineFuncer); ok {
			engineFuncer.AddFunc(funcName, funcBody)
		}
	}
}

// Load compiles all the registered engines.
func (v *View) Load() error {
	for i, n := 0, len(v.engines); i < n; i++ {
		e := v.engines[i]
		if err := e.Load(); err != nil {
			return err
		}
	}
	return nil
}
