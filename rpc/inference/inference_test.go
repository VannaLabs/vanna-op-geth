package inference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Basic inference test on a spread quoting regression model
func TestInference(t *testing.T) {
	rc := NewRequestClient(5125)
	tx := InferenceTx{
		Hash:   "0x123456789",
		Model:  "QmXQpupTphRTeXJMEz3BCt9YUF6kikcqExxPdcVoL1BBhy",
		Params: "[[0.002, 0.005, 0.004056685]]",
		TxType: "inference",
	}
	result, err := rc.Emit(tx)
	assert.Equal(t, nil, err)
	assert.Equal(t, result, "0.0013500629")
}

// Validate ECDSA Hex Signature
func TestSignatureValidation(t *testing.T) {
	engineNode := EngineNode{
		PublicKey:  "046fcc37ea5e9e09fec6c83e5fbd7a745e3eee81d16ebd861c9e66f55518c197984e9f113c07f875691df8afc1029496fc4cb9509b39dcd38f251a83359cc8b4f7",
		IPAddress:  "123.456.789",
		EthAddress: "0x123",
	}

	result := InferenceResult{
		Tx:    "0x123456789",
		Node:  "a35217ab3a12310813335af301368facb295a784bab8e542723f010071f38f68c5885713b6ab7a27491ce42ba3d1dd1d98c086bb216e80ba7b1622d401959f36",
		Value: "message",
	}
	valid, err := validateSignature(engineNode, result)
	assert.Equal(t, err, nil)
	assert.Equal(t, valid, true)

}
