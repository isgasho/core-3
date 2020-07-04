// Copyright 2020 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import (
	"crypto/ed25519"
	"crypto/x509"

	"v.io/v23/verror"
)

// NewED25519PublicKey creates a PublicKey object that uses the ED25519 algorithm and the provided ED25519 public key.
func NewED25519PublicKey(key ed25519.PublicKey) PublicKey {
	return newED25519PublicKeyImpl(key)
}

type ed25519PublicKey struct {
	key ed25519.PublicKey
}

func (pk *ed25519PublicKey) MarshalBinary() ([]byte, error) { return x509.MarshalPKIXPublicKey(pk.key) }
func (pk *ed25519PublicKey) String() string                 { return publicKeyString(pk) }
func (pk *ed25519PublicKey) verify(digest []byte, sig *Signature) bool {
	return ed25519.Verify(pk.key, digest, sig.Ed25519)
}

func (pk *ed25519PublicKey) hash() Hash {
	return SHA512Hash
}

// NewInMemoryED25519Signer creates a Signer that uses the provided ED25519
// private  key to sign messages.  This private key is kept in the clear in
// the memory of the running process.
func NewInMemoryED25519Signer(key ed25519.PrivateKey) Signer {
	// TODO(ashankar): Change this function to return an error
	// and not panic.
	signer, err := newInMemoryED25519SignerImpl(key)
	if err != nil {
		panic(err)
	}
	return signer
}

// NewED25519Signer creates a Signer that uses the provided function to sign
// messages.
func NewED25519Signer(key ed25519.PublicKey, sign func(data []byte) []byte) Signer {
	return &ed25519Signer{sign: sign, pubkey: NewED25519PublicKey(key)}
}

type ed25519Signer struct {
	sign   func(data []byte) []byte
	pubkey PublicKey
	// Object to hold on to for garbage collection
	impl interface{} //nolint:structcheck,unused
}

func (c *ed25519Signer) Sign(purpose, message []byte) (Signature, error) {
	hash := c.pubkey.hash()
	if message = messageDigest(hash, purpose, message, c.pubkey); message == nil {
		return Signature{}, verror.New(errSignCantHash, nil, hash)
	}
	return Signature{
		Purpose: purpose,
		Hash:    hash,
		Ed25519: c.sign(message),
	}, nil
}

func (c *ed25519Signer) PublicKey() PublicKey {
	return c.pubkey
}
