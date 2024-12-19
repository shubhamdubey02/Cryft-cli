// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package models

import (
	"fmt"

	"github.com/MetalBlockchain/metal-cli/pkg/constants"
	avago_constants "github.com/MetalBlockchain/metalgo/utils/constants"
)

type Network int64

const (
	Undefined Network = iota
	Mainnet
	Tahoe
	Local
)

func (s Network) String() string {
	switch s {
	case Mainnet:
		return "Mainnet"
	case Tahoe:
		return "Tahoe"
	case Local:
		return "Local Network"
	}
	return "Unknown Network"
}

func (s Network) NetworkID() (uint32, error) {
	switch s {
	case Mainnet:
		return avago_constants.MainnetID, nil
	case Tahoe:
		return avago_constants.TahoeID, nil
	case Local:
		return constants.LocalNetworkID, nil
	}
	return 0, fmt.Errorf("unsupported network")
}

func NetworkFromString(s string) Network {
	switch s {
	case Mainnet.String():
		return Mainnet
	case Tahoe.String():
		return Tahoe
	case Local.String():
		return Local
	}
	return Undefined
}

func NetworkFromNetworkID(networkID uint32) Network {
	switch networkID {
	case avago_constants.MainnetID:
		return Mainnet
	case avago_constants.TahoeID:
		return Tahoe
	case constants.LocalNetworkID:
		return Local
	}
	return Undefined
}
