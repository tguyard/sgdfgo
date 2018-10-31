package utils

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	assert := assert.New(t)

	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	encodeded, err := EncryptAes("DATA", key)
	assert.Equal(nil, err)
	decoded, err := DecryptAes(encodeded, key)
	assert.Equal(nil, err)
	assert.Equal("DATA", decoded)
}
