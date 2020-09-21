// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: flow

//nolint:golint
package flow

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

//////////////////////////////////////////////////
// Error definitions

var (
	ErrAuth          = verror.Register("v.io/v23/flow.Auth", verror.NoRetry, "{1:}{2:} {:3}")
	ErrNotTrusted    = verror.Register("v.io/v23/flow.NotTrusted", verror.NoRetry, "{1:}{2:} {:3}")
	ErrNetwork       = verror.Register("v.io/v23/flow.Network", verror.NoRetry, "{1:}{2:} {:3}")
	ErrDialFailed    = verror.Register("v.io/v23/flow.DialFailed", verror.NoRetry, "{1:}{2:} {:3}")
	ErrResolveFailed = verror.Register("v.io/v23/flow.ResolveFailed", verror.NoRetry, "{1:}{2:} {:3}")
	ErrProxy         = verror.Register("v.io/v23/flow.Proxy", verror.NoRetry, "{1:}{2:} {:3}")
	ErrBadArg        = verror.Register("v.io/v23/flow.BadArg", verror.NoRetry, "{1:}{2:} {:3}")
	ErrBadState      = verror.Register("v.io/v23/flow.BadState", verror.NoRetry, "{1:}{2:} {:3}")
	ErrAborted       = verror.Register("v.io/v23/flow.Aborted", verror.NoRetry, "{1:}{2:} {:3}")
)

// NewErrAuth returns an error with the ErrAuth ID.
func NewErrAuth(ctx *context.T, err error) error {
	return verror.New(ErrAuth, ctx, err)
}

// NewErrNotTrusted returns an error with the ErrNotTrusted ID.
func NewErrNotTrusted(ctx *context.T, err error) error {
	return verror.New(ErrNotTrusted, ctx, err)
}

// NewErrNetwork returns an error with the ErrNetwork ID.
func NewErrNetwork(ctx *context.T, err error) error {
	return verror.New(ErrNetwork, ctx, err)
}

// NewErrDialFailed returns an error with the ErrDialFailed ID.
func NewErrDialFailed(ctx *context.T, err error) error {
	return verror.New(ErrDialFailed, ctx, err)
}

// NewErrResolveFailed returns an error with the ErrResolveFailed ID.
func NewErrResolveFailed(ctx *context.T, err error) error {
	return verror.New(ErrResolveFailed, ctx, err)
}

// NewErrProxy returns an error with the ErrProxy ID.
func NewErrProxy(ctx *context.T, err error) error {
	return verror.New(ErrProxy, ctx, err)
}

// NewErrBadArg returns an error with the ErrBadArg ID.
func NewErrBadArg(ctx *context.T, err error) error {
	return verror.New(ErrBadArg, ctx, err)
}

// NewErrBadState returns an error with the ErrBadState ID.
func NewErrBadState(ctx *context.T, err error) error {
	return verror.New(ErrBadState, ctx, err)
}

// NewErrAborted returns an error with the ErrAborted ID.
func NewErrAborted(ctx *context.T, err error) error {
	return verror.New(ErrAborted, ctx, err)
}

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

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrAuth.ID), "{1:}{2:} {:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNotTrusted.ID), "{1:}{2:} {:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNetwork.ID), "{1:}{2:} {:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrDialFailed.ID), "{1:}{2:} {:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrResolveFailed.ID), "{1:}{2:} {:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrProxy.ID), "{1:}{2:} {:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBadArg.ID), "{1:}{2:} {:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBadState.ID), "{1:}{2:} {:3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrAborted.ID), "{1:}{2:} {:3}")

	return struct{}{}
}