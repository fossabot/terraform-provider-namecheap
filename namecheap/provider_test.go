package namecheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	assert.NoError(t, Provider().InternalValidate())
}
