package btcwire

import (
	"fmt"
	"io"
	"time"
)

// MsgAction represents a Gem game action.
// It is a gene
type MsgAction struct {
	// The version of this action message.
	// See MsgActionVersion.
	Version uint32

	Timestamp time.Time

	// The game this message is for.
	GameID ShaHash

	// The last message that the author saw before writing this message.
	// This could be all zeros if it is a game start message.
	PreviousMessageSignature [65]byte // could be 0s for game start

	// The player who is writing this message.
	Author [25]byte

	// The serialization of this message, signed.
	// When constructing and verifying this signature, assume these bytes
	// are set to 0.
	Signature [65]byte

	// The serialization of the game message.
	Message []byte
}

const (
	// MsgActionVersion is current action protocol version.
	MsgActionVersion = uint32(1)

	// minActionPayload is the minimum payload size for a Gem action.
	minActionPayload = 4 + 4 + 32 + 65 + 25 + 65 + 4 + 1

	// maxActionMessageLength is set low enough to prevent the length from
	// being large enough to DoS the receiver.
	maxActionMessageLength = 65535
)

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (a *MsgAction) BtcDecode(r io.Reader, pver uint32) error {
	err := readElement(r, &a.Version)
	if err != nil {
		return err
	}
	if a.Version != MsgActionVersion {
		return fmt.Errorf("Unknown action version %d", a.Version)
	}

	var dataLen uint32
	var sec uint32
	err = readElements(r, &sec, &a.GameID, &a.PreviousMessageSignature,
		&a.Author, &a.Signature, &dataLen)
	if err != nil {
		return err
	}
	a.Timestamp = time.Unix(int64(sec), 0)

	if dataLen > maxActionMessageLength {
		return fmt.Errorf("Action message length %d larger than maximum %d",
			dataLen, maxActionMessageLength)
	}

	a.Message = make([]byte, dataLen, dataLen)

	_, err = io.ReadFull(r, a.Message)
	if err != nil {
		return err
	}
	return nil
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (a *MsgAction) BtcEncode(w io.Writer, pver uint32) error {
	if len(a.Message) > maxActionMessageLength {
		return fmt.Errorf("Action message length %d larger than maximum %d",
			len(a.Message), maxActionMessageLength)
	}
	err := writeElements(w, a.Version, uint32(a.Timestamp.Unix()), a.GameID,
		a.PreviousMessageSignature, a.Author, a.Signature, uint32(len(a.Message)))
	if err != nil {
		return err
	}

	_, err = w.Write(a.Message)
	return err
}

// SerializeSize returns the number of bytes it would take to serialize the
// the action.
func (a *MsgAction) SerializeSize() int {
	return minActionPayload + len(a.Message) - 1
}
