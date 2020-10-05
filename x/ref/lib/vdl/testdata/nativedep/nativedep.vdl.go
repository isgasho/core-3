// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: nativedep

//nolint:golint
package nativedep

import (
	"time"

	"v.io/v23/vdl"
	"v.io/v23/vdl/testdata/nativetest"
	nativetest_2 "v.io/x/ref/lib/vdl/testdata/nativetest"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

// Type definitions
// ================

type All struct {
	A string
	B time.Time
	C nativetest.NativeSamePkg
	D map[nativetest.NativeSamePkg]time.Time
}

func (All) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativedep.All"`
}) {
}

func (x All) VDLIsZero() bool { //nolint:gocyclo
	if x.A != "" {
		return false
	}
	if !x.B.IsZero() {
		return false
	}
	if x.C != "" {
		return false
	}
	if x.D != nil {
		return false
	}
	return true
}

func (x All) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct1); err != nil {
		return err
	}
	if x.A != "" {
		var wire nativetest_2.WireString
		if err := nativetest_2.WireStringFromNative(&wire, x.A); err != nil {
			return err
		}
		if err := enc.NextFieldValueInt(0, vdlTypeInt322, int64(wire)); err != nil {
			return err
		}
	}
	if !x.B.IsZero() {
		var wire nativetest_2.WireTime
		if err := nativetest_2.WireTimeFromNative(&wire, x.B); err != nil {
			return err
		}
		if err := enc.NextFieldValueInt(1, vdlTypeInt323, int64(wire)); err != nil {
			return err
		}
	}
	if x.C != "" {
		var wire nativetest_2.WireSamePkg
		if err := nativetest_2.WireSamePkgFromNative(&wire, x.C); err != nil {
			return err
		}
		if err := enc.NextFieldValueInt(2, vdlTypeInt324, int64(wire)); err != nil {
			return err
		}
	}
	if x.D != nil {
		var wire nativetest_2.WireMultiImport
		if err := nativetest_2.WireMultiImportFromNative(&wire, x.D); err != nil {
			return err
		}
		if err := enc.NextFieldValueInt(3, vdlTypeInt325, int64(wire)); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *All) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = All{}
	if err := dec.StartValue(vdlTypeStruct1); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != vdlTypeStruct1 {
			index = vdlTypeStruct1.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			var wire nativetest_2.WireString
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := nativetest_2.WireStringToNative(wire, &x.A); err != nil {
				return err
			}
		case 1:
			var wire nativetest_2.WireTime
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := nativetest_2.WireTimeToNative(wire, &x.B); err != nil {
				return err
			}
		case 2:
			var wire nativetest_2.WireSamePkg
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := nativetest_2.WireSamePkgToNative(wire, &x.C); err != nil {
				return err
			}
		case 3:
			var wire nativetest_2.WireMultiImport
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := nativetest_2.WireMultiImportToNative(wire, &x.D); err != nil {
				return err
			}
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
//nolint:unused
var (
	vdlTypeStruct1 *vdl.Type
	vdlTypeInt322  *vdl.Type
	vdlTypeInt323  *vdl.Type
	vdlTypeInt324  *vdl.Type
	vdlTypeInt325  *vdl.Type
)

var initializeVDLCalled bool

// initializeVDL performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = initializeVDL()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func initializeVDL() struct{} {
	if initializeVDLCalled {
		return struct{}{}
	}
	initializeVDLCalled = true

	// Register types.
	vdl.Register((*All)(nil))

	// Initialize type definitions.
	vdlTypeStruct1 = vdl.TypeOf((*All)(nil)).Elem()
	vdlTypeInt322 = vdl.TypeOf((*nativetest_2.WireString)(nil))
	vdlTypeInt323 = vdl.TypeOf((*nativetest_2.WireTime)(nil))
	vdlTypeInt324 = vdl.TypeOf((*nativetest_2.WireSamePkg)(nil))
	vdlTypeInt325 = vdl.TypeOf((*nativetest_2.WireMultiImport)(nil))
	return struct{}{}
}
