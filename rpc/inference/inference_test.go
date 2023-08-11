package inference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Basic inference test on a volatility forecasting model
func TestInference(t *testing.T) {
	rc := NewRequestClient(5125)
	tx := InferenceTx{
		Hash:   "0x123",
		Model:  "Volatility",
		Params: "[[0.03],[0.05],[0.04056685],[0.03235871],[0.05629578]]",
	}
	result, err := rc.Emit(tx)
	assert.Equal(t, nil, err)
	assert.Equal(t, result, 0.053176194)
}
