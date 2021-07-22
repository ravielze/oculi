package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashWithCost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result, err := NewHashWithCost(MinCost + 2)
		assert.Nil(t, err)
		assert.Equal(t, &bcrypt{cost: MinCost + 2}, result)
	})

	t.Run("when less than min cost", func(t *testing.T) {
		result, err := NewHashWithCost(MinCost - 2)
		assert.Nil(t, err)
		assert.Equal(t, &bcrypt{cost: DefaultCost}, result)
	})

	t.Run("when higher than max cost", func(t *testing.T) {
		result, err := NewHashWithCost(MaxCost + 1)
		assert.Error(t, err)
		assert.Equal(t, ErrBcryptInvalidCost, err)
		assert.Nil(t, result)
	})
}

func Test_bcrypt_Hash(t *testing.T) {
	var (
		password = "somekind_of_password"
	)

	t.Run("success", func(t *testing.T) {
		b := &bcrypt{cost: DefaultCost}
		result, err := b.Hash(password)
		assert.Nil(t, err)
		assert.NotEqual(t, "", result)
	})
}

func Test_bcrypt_Verify(t *testing.T) {
	var (
		password       = "somekind_of_password"
		hashedPassword = "$2a$07$ASUgNfJjc3XuuqzqwIraWuh1KucN0i9TVPBgD9jRLmIFizoVqTqVS"
		wrongPassword  = "anykind_of_password"
		b              = &bcrypt{cost: MinCost + 3}
	)
	t.Run("password match", func(t *testing.T) {
		err := b.Verify(password, hashedPassword)
		assert.Nil(t, err)
	})

	t.Run("password mismatch", func(t *testing.T) {
		err := b.Verify(wrongPassword, hashedPassword)
		assert.Error(t, err)
		assert.Equal(t, err, ErrPasswordMismatch)
	})
}
