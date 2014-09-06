package btcwire

// MsgAction represents a Gem game action.
// It is a gene
type MsgAction struct {
	// The version of this action message.
	// See MsgActionVersion.
	Version uint32

	// The game this message is for.
	GameID ShaHash

	// The last message that the author saw before writing this message.
	// This could be all zeros if it is a game start message.
	PreviousMessageSignature []byte // could be 0s for game start

	// The player who is writing this message.
	Author ShaHash

	// The serialization of this message, signed.
	// When constructing and verifying this signature, assume these bytes
	// are set to 0.
	Signature []byte

	// The serialization of the game message.
	Message []byte
}

// The current action protocol version.
const MsgActionVersion = uint32(1)
