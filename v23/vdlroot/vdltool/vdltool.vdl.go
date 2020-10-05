// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vdltool

// Package vdltool defines types used by the vdl tool itself, including the
// format of vdl.config files.
//nolint:golint
package vdltool

import (
	"fmt"

	"v.io/v23/vdl"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

// Type definitions
// ================

// GenLanguage enumerates the known code generation languages.
type GenLanguage int

const (
	GenLanguageGo GenLanguage = iota
	GenLanguageJava
	GenLanguageJavascript
	GenLanguageSwift
)

// GenLanguageAll holds all labels for GenLanguage.
var GenLanguageAll = [...]GenLanguage{GenLanguageGo, GenLanguageJava, GenLanguageJavascript, GenLanguageSwift}

// GenLanguageFromString creates a GenLanguage from a string label.
//nolint:deadcode,unused
func GenLanguageFromString(label string) (x GenLanguage, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *GenLanguage) Set(label string) error {
	switch label {
	case "Go", "go":
		*x = GenLanguageGo
		return nil
	case "Java", "java":
		*x = GenLanguageJava
		return nil
	case "Javascript", "javascript":
		*x = GenLanguageJavascript
		return nil
	case "Swift", "swift":
		*x = GenLanguageSwift
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vdltool.GenLanguage", label)
}

// String returns the string label of x.
func (x GenLanguage) String() string {
	switch x {
	case GenLanguageGo:
		return "Go"
	case GenLanguageJava:
		return "Java"
	case GenLanguageJavascript:
		return "Javascript"
	case GenLanguageSwift:
		return "Swift"
	}
	return ""
}

func (GenLanguage) VDLReflect(struct {
	Name string `vdl:"vdltool.GenLanguage"`
	Enum struct{ Go, Java, Javascript, Swift string }
}) {
}

func (x GenLanguage) VDLIsZero() bool { //nolint:gocyclo
	return x == GenLanguageGo
}

func (x GenLanguage) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.WriteValueString(vdlTypeEnum1, x.String()); err != nil {
		return err
	}
	return nil
}

func (x *GenLanguage) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	switch value, err := dec.ReadValueString(); {
	case err != nil:
		return err
	default:
		if err := x.Set(value); err != nil {
			return err
		}
	}
	return nil
}

// GoKind describes the kind of Go type.
type GoKind int

const (
	GoKindStruct GoKind = iota
	GoKindBool
	GoKindNumber
	GoKindString
	GoKindArray
	GoKindSlice
	GoKindMap
	GoKindPointer
	GoKindIface
)

// GoKindAll holds all labels for GoKind.
var GoKindAll = [...]GoKind{GoKindStruct, GoKindBool, GoKindNumber, GoKindString, GoKindArray, GoKindSlice, GoKindMap, GoKindPointer, GoKindIface}

// GoKindFromString creates a GoKind from a string label.
//nolint:deadcode,unused
func GoKindFromString(label string) (x GoKind, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *GoKind) Set(label string) error {
	switch label {
	case "Struct", "struct":
		*x = GoKindStruct
		return nil
	case "Bool", "bool":
		*x = GoKindBool
		return nil
	case "Number", "number":
		*x = GoKindNumber
		return nil
	case "String", "string":
		*x = GoKindString
		return nil
	case "Array", "array":
		*x = GoKindArray
		return nil
	case "Slice", "slice":
		*x = GoKindSlice
		return nil
	case "Map", "map":
		*x = GoKindMap
		return nil
	case "Pointer", "pointer":
		*x = GoKindPointer
		return nil
	case "Iface", "iface":
		*x = GoKindIface
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vdltool.GoKind", label)
}

// String returns the string label of x.
func (x GoKind) String() string {
	switch x {
	case GoKindStruct:
		return "Struct"
	case GoKindBool:
		return "Bool"
	case GoKindNumber:
		return "Number"
	case GoKindString:
		return "String"
	case GoKindArray:
		return "Array"
	case GoKindSlice:
		return "Slice"
	case GoKindMap:
		return "Map"
	case GoKindPointer:
		return "Pointer"
	case GoKindIface:
		return "Iface"
	}
	return ""
}

func (GoKind) VDLReflect(struct {
	Name string `vdl:"vdltool.GoKind"`
	Enum struct{ Struct, Bool, Number, String, Array, Slice, Map, Pointer, Iface string }
}) {
}

func (x GoKind) VDLIsZero() bool { //nolint:gocyclo
	return x == GoKindStruct
}

func (x GoKind) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.WriteValueString(vdlTypeEnum2, x.String()); err != nil {
		return err
	}
	return nil
}

func (x *GoKind) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	switch value, err := dec.ReadValueString(); {
	case err != nil:
		return err
	default:
		if err := x.Set(value); err != nil {
			return err
		}
	}
	return nil
}

// GoZeroMode describes the relationship between the Go zero value of the native
// type, and the VDL zero value.
type GoZeroMode int

const (
	GoZeroModeUnknown GoZeroMode = iota
	GoZeroModeCanonical
	GoZeroModeUnique
)

// GoZeroModeAll holds all labels for GoZeroMode.
var GoZeroModeAll = [...]GoZeroMode{GoZeroModeUnknown, GoZeroModeCanonical, GoZeroModeUnique}

// GoZeroModeFromString creates a GoZeroMode from a string label.
//nolint:deadcode,unused
func GoZeroModeFromString(label string) (x GoZeroMode, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *GoZeroMode) Set(label string) error {
	switch label {
	case "Unknown", "unknown":
		*x = GoZeroModeUnknown
		return nil
	case "Canonical", "canonical":
		*x = GoZeroModeCanonical
		return nil
	case "Unique", "unique":
		*x = GoZeroModeUnique
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vdltool.GoZeroMode", label)
}

// String returns the string label of x.
func (x GoZeroMode) String() string {
	switch x {
	case GoZeroModeUnknown:
		return "Unknown"
	case GoZeroModeCanonical:
		return "Canonical"
	case GoZeroModeUnique:
		return "Unique"
	}
	return ""
}

func (GoZeroMode) VDLReflect(struct {
	Name string `vdl:"vdltool.GoZeroMode"`
	Enum struct{ Unknown, Canonical, Unique string }
}) {
}

func (x GoZeroMode) VDLIsZero() bool { //nolint:gocyclo
	return x == GoZeroModeUnknown
}

func (x GoZeroMode) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.WriteValueString(vdlTypeEnum3, x.String()); err != nil {
		return err
	}
	return nil
}

func (x *GoZeroMode) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	switch value, err := dec.ReadValueString(); {
	case err != nil:
		return err
	default:
		if err := x.Set(value); err != nil {
			return err
		}
	}
	return nil
}

// GoZero describes Go zero value behavior.
//
// REQUIRED: Either Mode == Unique or IsZero is set.  We will not perform
// native/wire conversions to check zero values.
type GoZero struct {
	Mode GoZeroMode
	// IsZero specifies a field, method or function that returns true iff the
	// native value represents the VDL zero value.
	//
	// If IsZero starts with a dot (.), it is assumed to be a field or method.
	// The field type or method return type must be bool.  Generated code will
	// apply the IsZero string verbatim to expressions of the native type.
	//
	// If IsZero doesn't start with a dot(.), it is assumed to be a function whose
	// return type must be bool.  Generated code will call the function with a
	// single argument that is an expression of the native type.
	//
	// TODO(toddw): The function form of IsZero isn't implemented.
	IsZero string
}

func (GoZero) VDLReflect(struct {
	Name string `vdl:"vdltool.GoZero"`
}) {
}

func (x GoZero) VDLIsZero() bool { //nolint:gocyclo
	return x == GoZero{}
}

func (x GoZero) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct4); err != nil {
		return err
	}
	if x.Mode != GoZeroModeUnknown {
		if err := enc.NextFieldValueString(0, vdlTypeEnum3, x.Mode.String()); err != nil {
			return err
		}
	}
	if x.IsZero != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.IsZero); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *GoZero) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = GoZero{}
	if err := dec.StartValue(vdlTypeStruct4); err != nil {
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
		if decType != vdlTypeStruct4 {
			index = vdlTypeStruct4.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				if err := x.Mode.Set(value); err != nil {
					return err
				}
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.IsZero = value
			}
		}
	}
}

// GoImport describes Go import information.
type GoImport struct {
	// Path is the package path that uniquely identifies the imported package.
	Path string
	// Name is the name of the package identified by Path.  Due to Go conventions,
	// it is typically just the basename of Path, but may be set to something
	// different if the imported package doesn't follow Go conventions.
	Name string
}

func (GoImport) VDLReflect(struct {
	Name string `vdl:"vdltool.GoImport"`
}) {
}

func (x GoImport) VDLIsZero() bool { //nolint:gocyclo
	return x == GoImport{}
}

func (x GoImport) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct5); err != nil {
		return err
	}
	if x.Path != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Path); err != nil {
			return err
		}
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *GoImport) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = GoImport{}
	if err := dec.StartValue(vdlTypeStruct5); err != nil {
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
		if decType != vdlTypeStruct5 {
			index = vdlTypeStruct5.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Path = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		}
	}
}

// GoType describes the Go type information associated with a VDL type.
// See v.io/x/ref/lib/vdl/testdata/native for examples.
type GoType struct {
	// Kind is the kind of Type.
	Kind GoKind
	// Type is the Go type to use in generated code, instead of the VDL type.
	Type string
	// Zero specifies zero value setting and checking behavior.
	Zero GoZero
	// ToNative and FromNative overrides the names of the native conversion
	// functions.  If unspecified, the default functions are of the form:
	//   WireTypeToNative(wire WireType, native *NativeType) error
	//   WireTypeFromNative(wire *WireType, native NativeType) error
	ToNative   string
	FromNative string
	// Imports are the Go imports referenced by the Type, ToNative and FromNative
	// fields.  In each field, use the standard local package name, and specify
	// the import package here.  E.g. to specify the native type time.Time:
	//   GoType{
	//     Kind:    Struct,
	//     Type:    "time.Time",
	//     Zero:    {Mode: Canonical, IsZero: ".IsZero()"},
	//     Imports: {{Path: "time", Name: "time"}},
	//   }
	Imports []GoImport
}

func (GoType) VDLReflect(struct {
	Name string `vdl:"vdltool.GoType"`
}) {
}

func (x GoType) VDLIsZero() bool { //nolint:gocyclo
	if x.Kind != GoKindStruct {
		return false
	}
	if x.Type != "" {
		return false
	}
	if x.Zero != (GoZero{}) {
		return false
	}
	if x.ToNative != "" {
		return false
	}
	if x.FromNative != "" {
		return false
	}
	if len(x.Imports) != 0 {
		return false
	}
	return true
}

func (x GoType) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct6); err != nil {
		return err
	}
	if x.Kind != GoKindStruct {
		if err := enc.NextFieldValueString(0, vdlTypeEnum2, x.Kind.String()); err != nil {
			return err
		}
	}
	if x.Type != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.Type); err != nil {
			return err
		}
	}
	if x.Zero != (GoZero{}) {
		if err := enc.NextField(2); err != nil {
			return err
		}
		if err := x.Zero.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.ToNative != "" {
		if err := enc.NextFieldValueString(3, vdl.StringType, x.ToNative); err != nil {
			return err
		}
	}
	if x.FromNative != "" {
		if err := enc.NextFieldValueString(4, vdl.StringType, x.FromNative); err != nil {
			return err
		}
	}
	if len(x.Imports) != 0 {
		if err := enc.NextField(5); err != nil {
			return err
		}
		if err := vdlWriteAnonList1(enc, x.Imports); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func vdlWriteAnonList1(enc vdl.Encoder, x []GoImport) error {
	if err := enc.StartValue(vdlTypeList7); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *GoType) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = GoType{}
	if err := dec.StartValue(vdlTypeStruct6); err != nil {
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
		if decType != vdlTypeStruct6 {
			index = vdlTypeStruct6.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				if err := x.Kind.Set(value); err != nil {
					return err
				}
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Type = value
			}
		case 2:
			if err := x.Zero.VDLRead(dec); err != nil {
				return err
			}
		case 3:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.ToNative = value
			}
		case 4:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.FromNative = value
			}
		case 5:
			if err := vdlReadAnonList1(dec, &x.Imports); err != nil {
				return err
			}
		}
	}
}

func vdlReadAnonList1(dec vdl.Decoder, x *[]GoImport) error {
	if err := dec.StartValue(vdlTypeList7); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]GoImport, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem GoImport
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

// GoConfig specifies go specific configuration.
type GoConfig struct {
	// WireToNativeTypes specifies the mapping from a VDL wire type to its Go
	// native type representation.  This is rarely used and easy to configure
	// incorrectly; usage is currently restricted to packages that are explicitly
	// whitelisted.
	//
	// WireToNativeTypes are meant for scenarios where there is an idiomatic Go
	// type used in your code, but you need a standard VDL representation for wire
	// compatibility.  E.g. the VDL time package defines Duration and Time for
	// wire compatibility, but we want the generated code to use the standard Go
	// time package.
	//
	// The key of the map is the name of the VDL type (aka WireType), which must
	// be defined in the vdl package associated with the vdl.config file.
	//
	// The code generator assumes the existence of a pair of conversion functions
	// converting between the wire and native types, and will automatically call
	// vdl.RegisterNative with these function names.
	//
	// Assuming the name of the WireType is Foo:
	//   func fooToNative(x Foo, n *Native) error
	//   func fooFromNative(x *Foo, n Native) error
	WireToNativeTypes map[string]GoType
}

func (GoConfig) VDLReflect(struct {
	Name string `vdl:"vdltool.GoConfig"`
}) {
}

func (x GoConfig) VDLIsZero() bool { //nolint:gocyclo
	return len(x.WireToNativeTypes) == 0
}

func (x GoConfig) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct8); err != nil {
		return err
	}
	if len(x.WireToNativeTypes) != 0 {
		if err := enc.NextField(0); err != nil {
			return err
		}
		if err := vdlWriteAnonMap2(enc, x.WireToNativeTypes); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func vdlWriteAnonMap2(enc vdl.Encoder, x map[string]GoType) error {
	if err := enc.StartValue(vdlTypeMap9); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, key); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *GoConfig) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = GoConfig{}
	if err := dec.StartValue(vdlTypeStruct8); err != nil {
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
		if decType != vdlTypeStruct8 {
			index = vdlTypeStruct8.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		if index == 0 {

			if err := vdlReadAnonMap2(dec, &x.WireToNativeTypes); err != nil {
				return err
			}
		}
	}
}

func vdlReadAnonMap2(dec vdl.Decoder, x *map[string]GoType) error {
	if err := dec.StartValue(vdlTypeMap9); err != nil {
		return err
	}
	var tmpMap map[string]GoType
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[string]GoType, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			var elem GoType
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			if tmpMap == nil {
				tmpMap = make(map[string]GoType)
			}
			tmpMap[key] = elem
		}
	}
}

// JavaConfig specifies java specific configuration.
type JavaConfig struct {
	// WireToNativeTypes specifies the mapping from a VDL wire type to its Java
	// native type representation.  This is rarely used and easy to configure
	// incorrectly; usage is currently restricted to packages that are explicitly
	// whitelisted.
	//
	// WireToNativeTypes are meant for scenarios where there is an idiomatic Java
	// type used in your code, but you need a standard VDL representation for wire
	// compatibility.  E.g. the VDL time package defines Duration and Time for
	// wire compatibility, but we want the generated code to use the org.joda.time
	// package.
	//
	// The key of the map is the name of the VDL type (aka WireType), which must
	// be defined in the vdl package associated with the vdl.config file.
	//
	// The code generator assumes that the conversion functions will be registered
	// in java vdl package.
	WireToNativeTypes map[string]string
	// WireTypeRenames specifies the mapping from a VDL wire type name to its
	// Java native type name.
	//
	// WireTypeRenames are meant for scenarios where the VDL wire name
	// conflicts in some way with the Java native names, e.g., a VDL Integer
	// type could be named VInteger for clarity.
	//
	// When combined with WireToNativeTypes, this feature allows us to attach
	// functions to VDL types.  For example, we may rename AccessList VDL type
	// into WireAccessList and then map WireAccessList to our Java native type
	// AccessList which defines functions on the VDL data.
	//
	// The key of the map is the name of the VDL wire type, which must be
	// defined in the vdl package associated with the vdl.config file.
	WireTypeRenames map[string]string
}

func (JavaConfig) VDLReflect(struct {
	Name string `vdl:"vdltool.JavaConfig"`
}) {
}

func (x JavaConfig) VDLIsZero() bool { //nolint:gocyclo
	if len(x.WireToNativeTypes) != 0 {
		return false
	}
	if len(x.WireTypeRenames) != 0 {
		return false
	}
	return true
}

func (x JavaConfig) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct10); err != nil {
		return err
	}
	if len(x.WireToNativeTypes) != 0 {
		if err := enc.NextField(0); err != nil {
			return err
		}
		if err := vdlWriteAnonMap3(enc, x.WireToNativeTypes); err != nil {
			return err
		}
	}
	if len(x.WireTypeRenames) != 0 {
		if err := enc.NextField(1); err != nil {
			return err
		}
		if err := vdlWriteAnonMap3(enc, x.WireTypeRenames); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func vdlWriteAnonMap3(enc vdl.Encoder, x map[string]string) error {
	if err := enc.StartValue(vdlTypeMap11); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, key); err != nil {
			return err
		}
		if err := enc.WriteValueString(vdl.StringType, elem); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *JavaConfig) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = JavaConfig{}
	if err := dec.StartValue(vdlTypeStruct10); err != nil {
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
		if decType != vdlTypeStruct10 {
			index = vdlTypeStruct10.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			if err := vdlReadAnonMap3(dec, &x.WireToNativeTypes); err != nil {
				return err
			}
		case 1:
			if err := vdlReadAnonMap3(dec, &x.WireTypeRenames); err != nil {
				return err
			}
		}
	}
}

func vdlReadAnonMap3(dec vdl.Decoder, x *map[string]string) error {
	if err := dec.StartValue(vdlTypeMap11); err != nil {
		return err
	}
	var tmpMap map[string]string
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[string]string, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			var elem string
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				elem = value
			}
			if tmpMap == nil {
				tmpMap = make(map[string]string)
			}
			tmpMap[key] = elem
		}
	}
}

// JavascriptConfig specifies javascript specific configuration.
type JavascriptConfig struct {
}

func (JavascriptConfig) VDLReflect(struct {
	Name string `vdl:"vdltool.JavascriptConfig"`
}) {
}

func (x JavascriptConfig) VDLIsZero() bool { //nolint:gocyclo
	return x == JavascriptConfig{}
}

func (x JavascriptConfig) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct12); err != nil {
		return err
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *JavascriptConfig) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = JavascriptConfig{}
	if err := dec.StartValue(vdlTypeStruct12); err != nil {
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
		if decType != vdlTypeStruct12 {
			index = vdlTypeStruct12.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		}
	}
}

// SwiftConfig specifies swift specific configuration for this package.
// Note that despite the SwiftConfig options for a given package (which should be
// very rare in practice), we still need to know the name of the swift module
// that this package relates to to properly understand import boundaries between
// projects/frameworks/modules.
//
// We do this by defining a file called "swiftmodule" that contains JUST the
// name of the Swift module at the root of your VDL packages. For example,
// if you have the VDL files for your Xcode project/target called UberForCats
// at /Users/aaron/uberforcats/vdl, then create
// /Users/aaron/uberforcats/vdl/com.uberforcats/swiftmodule and have it just contain
// "UberForCats". We then will treat any VDL files contained in that directory and
// any subdirectories as part of the UberForCats Swift module, ultimately letting
// the compiler and will automatically do the right thing if others import your package.
// If you don't do this then nobody will be able to import your VDL types in Swift,
// and you might end up with extra long class/pkg names (ComuberforcatsServicesProfit
// instead of ServicesProfit for $VDLROOT/com.uberforcats/services/profit).
//
// If you are creating multiple Swift modules for a given $VDLROOT then just place
// swiftmodule files at the logical boundaries. For eample, we do this for v.io/v23
// to be exported to the VanadiumCore framework, but everything under v.io/v23/services
// lives in the VanadiumServices framework.
type SwiftConfig struct {
	// WireToNativeTypes specifies the mapping from a VDL wire type to its Swift
	// native type representation.  This is rarely used and easy to configure
	// incorrectly; usage is currently restricted to packages that are explicitly
	// whitelisted.
	//
	// WireToNativeTypes are meant for scenarios where there is an idiomatic Swift
	// type used in your code, but you need a standard VDL representation for wire
	// compatibility.  E.g. the VDL time package defines Duration and Time for
	// wire compatibility, but we want the generated code to use NSDate or NSTimeInterval
	//
	// The key of the map is the name of the VDL type (aka WireType), which must
	// be defined in the vdl package associated with the vdl.config file.
	//
	// The code generator assumes that the conversion functions will be registered
	// in Swift vdl package.
	WireToNativeTypes map[string]string
}

func (SwiftConfig) VDLReflect(struct {
	Name string `vdl:"vdltool.SwiftConfig"`
}) {
}

func (x SwiftConfig) VDLIsZero() bool { //nolint:gocyclo
	return len(x.WireToNativeTypes) == 0
}

func (x SwiftConfig) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct13); err != nil {
		return err
	}
	if len(x.WireToNativeTypes) != 0 {
		if err := enc.NextField(0); err != nil {
			return err
		}
		if err := vdlWriteAnonMap3(enc, x.WireToNativeTypes); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *SwiftConfig) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = SwiftConfig{}
	if err := dec.StartValue(vdlTypeStruct13); err != nil {
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
		if decType != vdlTypeStruct13 {
			index = vdlTypeStruct13.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		if index == 0 {

			if err := vdlReadAnonMap3(dec, &x.WireToNativeTypes); err != nil {
				return err
			}
		}
	}
}

// Config specifies the configuration for the vdl tool.  This is typically
// represented in optional "vdl.config" files in each vdl source package.  Each
// vdl.config file implicitly imports this package.  E.g. you may refer to
// vdltool.Config in the "vdl.config" file without explicitly importing vdltool.
type Config struct {
	// GenLanguages restricts the set of code generation languages.  If the set is
	// empty, all supported languages are allowed to be generated.
	GenLanguages map[GenLanguage]struct{}
	// Language-specific configurations.
	Go         GoConfig
	Java       JavaConfig
	Javascript JavascriptConfig
	Swift      SwiftConfig
}

func (Config) VDLReflect(struct {
	Name string `vdl:"vdltool.Config"`
}) {
}

func (x Config) VDLIsZero() bool { //nolint:gocyclo
	if len(x.GenLanguages) != 0 {
		return false
	}
	if !x.Go.VDLIsZero() {
		return false
	}
	if !x.Java.VDLIsZero() {
		return false
	}
	if x.Javascript != (JavascriptConfig{}) {
		return false
	}
	if !x.Swift.VDLIsZero() {
		return false
	}
	return true
}

func (x Config) VDLWrite(enc vdl.Encoder) error { //nolint:gocyclo
	if err := enc.StartValue(vdlTypeStruct14); err != nil {
		return err
	}
	if len(x.GenLanguages) != 0 {
		if err := enc.NextField(0); err != nil {
			return err
		}
		if err := vdlWriteAnonSet4(enc, x.GenLanguages); err != nil {
			return err
		}
	}
	if !x.Go.VDLIsZero() {
		if err := enc.NextField(1); err != nil {
			return err
		}
		if err := x.Go.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.Java.VDLIsZero() {
		if err := enc.NextField(2); err != nil {
			return err
		}
		if err := x.Java.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Javascript != (JavascriptConfig{}) {
		if err := enc.NextField(3); err != nil {
			return err
		}
		if err := x.Javascript.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.Swift.VDLIsZero() {
		if err := enc.NextField(4); err != nil {
			return err
		}
		if err := x.Swift.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func vdlWriteAnonSet4(enc vdl.Encoder, x map[GenLanguage]struct{}) error {
	if err := enc.StartValue(vdlTypeSet15); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key := range x {
		if err := enc.NextEntryValueString(vdlTypeEnum1, key.String()); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Config) VDLRead(dec vdl.Decoder) error { //nolint:gocyclo
	*x = Config{}
	if err := dec.StartValue(vdlTypeStruct14); err != nil {
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
		if decType != vdlTypeStruct14 {
			index = vdlTypeStruct14.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			if err := vdlReadAnonSet4(dec, &x.GenLanguages); err != nil {
				return err
			}
		case 1:
			if err := x.Go.VDLRead(dec); err != nil {
				return err
			}
		case 2:
			if err := x.Java.VDLRead(dec); err != nil {
				return err
			}
		case 3:
			if err := x.Javascript.VDLRead(dec); err != nil {
				return err
			}
		case 4:
			if err := x.Swift.VDLRead(dec); err != nil {
				return err
			}
		}
	}
}

func vdlReadAnonSet4(dec vdl.Decoder, x *map[GenLanguage]struct{}) error {
	if err := dec.StartValue(vdlTypeSet15); err != nil {
		return err
	}
	var tmpMap map[GenLanguage]struct{}
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[GenLanguage]struct{}, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			var keyEnum GenLanguage
			if err := keyEnum.Set(key); err != nil {
				return err
			}
			if tmpMap == nil {
				tmpMap = make(map[GenLanguage]struct{})
			}
			tmpMap[keyEnum] = struct{}{}
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
//nolint:unused
var (
	vdlTypeEnum1    *vdl.Type
	vdlTypeEnum2    *vdl.Type
	vdlTypeEnum3    *vdl.Type
	vdlTypeStruct4  *vdl.Type
	vdlTypeStruct5  *vdl.Type
	vdlTypeStruct6  *vdl.Type
	vdlTypeList7    *vdl.Type
	vdlTypeStruct8  *vdl.Type
	vdlTypeMap9     *vdl.Type
	vdlTypeStruct10 *vdl.Type
	vdlTypeMap11    *vdl.Type
	vdlTypeStruct12 *vdl.Type
	vdlTypeStruct13 *vdl.Type
	vdlTypeStruct14 *vdl.Type
	vdlTypeSet15    *vdl.Type
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
	vdl.Register((*GenLanguage)(nil))
	vdl.Register((*GoKind)(nil))
	vdl.Register((*GoZeroMode)(nil))
	vdl.Register((*GoZero)(nil))
	vdl.Register((*GoImport)(nil))
	vdl.Register((*GoType)(nil))
	vdl.Register((*GoConfig)(nil))
	vdl.Register((*JavaConfig)(nil))
	vdl.Register((*JavascriptConfig)(nil))
	vdl.Register((*SwiftConfig)(nil))
	vdl.Register((*Config)(nil))

	// Initialize type definitions.
	vdlTypeEnum1 = vdl.TypeOf((*GenLanguage)(nil))
	vdlTypeEnum2 = vdl.TypeOf((*GoKind)(nil))
	vdlTypeEnum3 = vdl.TypeOf((*GoZeroMode)(nil))
	vdlTypeStruct4 = vdl.TypeOf((*GoZero)(nil)).Elem()
	vdlTypeStruct5 = vdl.TypeOf((*GoImport)(nil)).Elem()
	vdlTypeStruct6 = vdl.TypeOf((*GoType)(nil)).Elem()
	vdlTypeList7 = vdl.TypeOf((*[]GoImport)(nil))
	vdlTypeStruct8 = vdl.TypeOf((*GoConfig)(nil)).Elem()
	vdlTypeMap9 = vdl.TypeOf((*map[string]GoType)(nil))
	vdlTypeStruct10 = vdl.TypeOf((*JavaConfig)(nil)).Elem()
	vdlTypeMap11 = vdl.TypeOf((*map[string]string)(nil))
	vdlTypeStruct12 = vdl.TypeOf((*JavascriptConfig)(nil)).Elem()
	vdlTypeStruct13 = vdl.TypeOf((*SwiftConfig)(nil)).Elem()
	vdlTypeStruct14 = vdl.TypeOf((*Config)(nil)).Elem()
	vdlTypeSet15 = vdl.TypeOf((*map[GenLanguage]struct{})(nil))
	return struct{}{}
}
