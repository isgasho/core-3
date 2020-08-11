// Copyright 2020 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: ecdsaonly

//nolint:golint
package ecdsaonly

import (
	"v.io/v23/vdl"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

//////////////////////////////////////////////////
// Type definitions

type Hash string

func (Hash) VDLReflect(struct {
	Name string `vdl:"v.io/v23/security/internal/ecdsaonly.Hash"`
}) {
}

func (x Hash) VDLIsZero() bool { //nolint:gocyclo
	return x == ""
}

func (x Hash) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.WriteValueString(vdlTypeString1, string(x)); err != nil {
		return err
	}
	return nil
}

func (x *Hash) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	switch value, err := dec.ReadValueString(); {
	case err != nil:
		return err
	default:
		*x = Hash(value)
	}
	return nil
}

// Signature represents a digital signature.
type Signature struct {
	Purpose []byte
	Hash    Hash
	R       []byte
	S       []byte
}

func (Signature) VDLReflect(struct {
	Name string `vdl:"v.io/v23/security/internal/ecdsaonly.Signature"`
}) {
}

func (x Signature) VDLIsZero() bool { //nolint:gocyclo
	if len(x.Purpose) != 0 {
		return false
	}
	if x.Hash != "" {
		return false
	}
	if len(x.R) != 0 {
		return false
	}
	if len(x.S) != 0 {
		return false
	}
	return true
}

func (x Signature) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct2); err != nil {
		return err
	}
	if len(x.Purpose) != 0 {
		if err := enc.NextFieldValueBytes(0, vdlTypeList3, x.Purpose); err != nil {
			return err
		}
	}
	if x.Hash != "" {
		if err := enc.NextFieldValueString(1, vdlTypeString1, string(x.Hash)); err != nil {
			return err
		}
	}
	if len(x.R) != 0 {
		if err := enc.NextFieldValueBytes(2, vdlTypeList3, x.R); err != nil {
			return err
		}
	}
	if len(x.S) != 0 {
		if err := enc.NextFieldValueBytes(3, vdlTypeList3, x.S); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Signature) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = Signature{}
	if err := dec.StartValue(vdlTypeStruct2); err != nil {
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
		if decType != vdlTypeStruct2 {
			index = vdlTypeStruct2.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			if err := dec.ReadValueBytes(-1, &x.Purpose); err != nil {
				return err
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Hash = Hash(value)
			}
		case 2:
			if err := dec.ReadValueBytes(-1, &x.R); err != nil {
				return err
			}
		case 3:
			if err := dec.ReadValueBytes(-1, &x.S); err != nil {
				return err
			}
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
//nolint:unused
var (
	vdlTypeString1 *vdl.Type
	vdlTypeStruct2 *vdl.Type
	vdlTypeList3   *vdl.Type
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
	vdl.Register((*Hash)(nil))
	vdl.Register((*Signature)(nil))

	// Initialize type definitions.
	vdlTypeString1 = vdl.TypeOf((*Hash)(nil))
	vdlTypeStruct2 = vdl.TypeOf((*Signature)(nil)).Elem()
	vdlTypeList3 = vdl.TypeOf((*[]byte)(nil))

	return struct{}{}
}
