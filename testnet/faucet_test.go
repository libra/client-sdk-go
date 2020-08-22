// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet_test

import (
	"testing"

	"github.com/libra/libra-client-sdk-go/librakeys"
	"github.com/libra/libra-client-sdk-go/testnet"
	"github.com/stretchr/testify/assert"
)

func TestMint(t *testing.T) {
	keys := librakeys.MustGenKeys()
	testnet.MustMint(keys.AuthKey.Hex(), 1000, "LBR")
}

func TestMustMintPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
		assert.Fail(t, "should panic, but not")
	}()

	testnet.MustMint("invalid", 1000, "HELLO")
}
