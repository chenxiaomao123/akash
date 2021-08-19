package keeper

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/types/address"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ovrclk/akash/x/audit/types"
)

func providerKey(id types.ProviderID) []byte {
	buf := bytes.NewBuffer(types.PrefixProviderID())
	if _, err := buf.Write(address.MustLengthPrefix(id.Owner.Bytes())); err != nil {
		panic(err)
	}

	if _, err := buf.Write(address.MustLengthPrefix(id.Auditor.Bytes())); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func providerPrefix(id sdk.Address) []byte {
	buf := bytes.NewBuffer(types.PrefixProviderID())
	if _, err := buf.Write(address.MustLengthPrefix(id.Bytes())); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
