package mode

// ModeType is the encrypted mode.
type ModeType uint8

const (
	// CBC mode, each block of plaintext is XORed with the previous ciphertext block before being encrypted.
	// This way, each ciphertext block depends on all plaintext blocks processed up to that point.
	// To make each message unique, an initialization vector must be used in the first block.
	CBC ModeType = 1 + iota

	// CFB mode, in its simplest form uses the entire output of the block cipher. In this
	// variation, it is very similar to CBC, makes a block cipher into a self-synchronizing stream cipher.
	CFB

	// OFB (short for output feedback) is an AES block cipher mode similar to the CFB mode.
	// What mainly differs from CFB is that the OFB mode relies on XOR-ing plaintext and ciphertext blocks
	// with expanded versions of the initialization vector.
	OFB

	// CTR is a Stream which encrypts/decrypts using the given Block in
	// counter mode. The length of iv must be the same as the Block's block size.
	CTR
)
