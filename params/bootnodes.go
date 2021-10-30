// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/citrino-network/citrino-cli/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://e700ba1de15a6af3c38f45f60324018b7abd99117ce141a48793e2e479707a0ffa70a63a615bf0c5754760fb7cd2470d2059ac3b2e77cfcd61b8cf275372b35b@181.129.103.142:30303",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
}

var V5Bootnodes = []string{
	"enr:-KO4QHDYX79GOtbwGRZHP92bNKftJ9bpRrw-f7Fa67xrROrMDJJvGjP5vXL29ZJIZD_oWI2a0-SUbr7SAyzr5bZUTTWGAXzSL6lDg2V0aMfGhAL0oZ6AgmlkgnY0gmlwhLWBZ46Jc2VjcDI1NmsxoQPnALod4Vpq88OPRfYDJAGLer2ZEXzhQaSHk-LkeXB6D4RzbmFwwIN0Y3CCdl-DdWRwgnZf",
}

const dnsPrefix = "Nil@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet" //
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
