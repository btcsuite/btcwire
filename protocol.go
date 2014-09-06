// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcwire

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ProtocolVersion is the latest protocol version this package supports.
	ProtocolVersion uint32 = 70003

	// MultipleAddressVersion is the protocol version which added multiple
	// addresses per message (pver >= MultipleAddressVersion).
	MultipleAddressVersion uint32 = 209

	// NetAddressTimeVersion is the protocol version which added the
	// timestamp field (pver >= NetAddressTimeVersion).
	NetAddressTimeVersion uint32 = 31402

	// BIP0031Version is the protocol version AFTER which a pong message
	// and nonce field in ping were added (pver > BIP0031Version).
	BIP0031Version uint32 = 60000

	// BIP0035Version is the protocol version which added the mempool
	// message (pver >= BIP0035Version).
	BIP0035Version uint32 = 60002

	// BIP0037Version is the protocol version which added new connection
	// bloom filtering related messages and extended the version message
	// with a relay flag (pver >= BIP0037Version).
	BIP0037Version uint32 = 70001

	// RejectVersion is the protocol version which added a new reject
	// message.
	RejectVersion uint32 = 70002
)

// ServiceFlag identifies services supported by a bitcoin peer.
type ServiceFlag uint64

const (
	// SFNodeNetwork is a flag used to indicate a peer is a full node.
	SFNodeNetwork ServiceFlag = 1 << iota
)

// Map of service flags back to their constant names for pretty printing.
var sfStrings = map[ServiceFlag]string{
	SFNodeNetwork: "SFNodeNetwork",
}

// String returns the ServiceFlag in human-readable form.
func (f ServiceFlag) String() string {
	// No flags are set.
	if f == 0 {
		return "0x0"
	}

	// Add individual bit flags.
	s := ""
	for flag, name := range sfStrings {
		if f&flag == flag {
			s += name + "|"
			f -= flag
		}
	}

	// Add any remaining flags which aren't accounted for as hex.
	s = strings.TrimRight(s, "|")
	if f != 0 {
		s += "|0x" + strconv.FormatUint(uint64(f), 16)
	}
	s = strings.TrimLeft(s, "|")
	return s
}

// BitcoinNet represents which bitcoin network a message belongs to.
type BitcoinNet uint32

// Constants used to indicate the message bitcoin network.  They can also be
// used to seek to the next message when a stream's state is unknown, but
// this package does not provide that functionality since it's generally a
// better idea to simply disconnect clients that are misbehaving over TCP.
const (
	// TestNet represents the Playcoin test network.
	TestNet BitcoinNet = 0x9a25c7b3
)

// bnStrings is a map of bitcoin networks back to their constant names for
// pretty printing.
var bnStrings = map[BitcoinNet]string{
	TestNet: "GemTestNet",
}

// String returns the BitcoinNet in human-readable form.
func (n BitcoinNet) String() string {
	if s, ok := bnStrings[n]; ok {
		return s
	}

	return fmt.Sprintf("Unknown BitcoinNet (%d)", uint32(n))
}
